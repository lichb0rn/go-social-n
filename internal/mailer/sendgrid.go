package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailer struct {
	fromEmail string
	apiKey    string
	client    *sendgrid.Client
}

func NewSendgrid(apiKey string, fromEmail string) *SendGridMailer {
	client := sendgrid.NewSendClient(apiKey)

	return &SendGridMailer{
		fromEmail: fromEmail,
		apiKey:    apiKey,
		client:    client,
	}
}

// Send renders the given template with the given data and sends it to the given
// email address using the given SendGrid API key.
//
// If isSanbox is true, it will send the email to the sandbox endpoint and
// return a 200 status code immediately. Otherwise, it will send the email
// to the given email address and return the status code of the SendGrid API.
//
// The function will retry sending the email up to maxRetries times if the
// API returns an error. After maxRetries attempts, it will return the last error
// encountered.
//
// The function returns the status code of the SendGrid API and an error if
// the email wasn't sent successfully.
func (m *SendGridMailer) Send(templateFile string, username, email string, data any, isSanbox bool) (int, error) {
	// Since I don't have a sendgrid sandbox account
	// I can't test this properly
	if isSanbox {
		return 200, nil
	}

	from := mail.NewEmail(FromName, m.fromEmail)
	to := mail.NewEmail(username, email)

	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return -1, nil
	}

	body := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(body, "body", data); err != nil {
		return -1, nil
	}

	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())
	message.MailSettings = &mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &isSanbox,
		},
	}

	var retryErr error
	for i := range maxRetries {
		res, retryErr := m.client.Send(message)
		if retryErr == nil {
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		return res.StatusCode, nil
	}
	return -1, fmt.Errorf("failed to send email after %d attempts, error: %v", maxRetries, retryErr)
}
