package main

import (
	"net/http"
	"net/http/httptest"
	"social/internal/ratelimiter"
	"testing"
	"time"
)

func TestRateLimiterMiddleware(t *testing.T) {
	cfg := config{
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: 20,
			TimeFrame:            time.Second * 5,
			Enabled:              true,
		},
		addr: ":8080",
	}

	app := newTestApplication(t, cfg)
	ts := httptest.NewServer(app.mount())
	defer ts.Close()

	client := &http.Client{}
	mockIP := "192.168.0.1"
	marginOfError := 2

	for i := range cfg.rateLimiter.RequestsPerTimeFrame + marginOfError {
		req, err := http.NewRequest("GET", ts.URL+"/v1/health", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		req.Header.Set("X-Forwarded-For", mockIP)

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if i < cfg.rateLimiter.RequestsPerTimeFrame {
			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
			}
		} else {
			if resp.StatusCode != http.StatusTooManyRequests {
				t.Errorf("expected status code %d, got %d", http.StatusTooManyRequests, resp.StatusCode)
			}
		}
	}
}
