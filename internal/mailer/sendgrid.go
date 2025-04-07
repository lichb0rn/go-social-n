package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
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

func (m *SendGridMailer) Send(templateFile string, username, email string, data any, isSanbox bool) error {
	// Since I don't have a sendgrid sandbox account
	// I can't test this properly
	if isSanbox {
		return nil
	}

	from := mail.NewEmail(FromName, m.fromEmail)
	to := mail.NewEmail(username, email)

	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return nil
	}

	body := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(body, "body", data); err != nil {
		return nil
	}

	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())
	message.MailSettings = &mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &isSanbox,
		},
	}

	for i := range maxRetries {
		res, err := m.client.Send(message)
		if err == nil {
			log.Printf("failed to send email %v, attempt %d of %d", email, i+1, maxRetries)
			log.Printf("Error: %v", err)

			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		log.Printf("Email sent with status code %v", res.StatusCode)
		return nil
	}
	return fmt.Errorf("failed to send email after %d attempts", maxRetries)
}
