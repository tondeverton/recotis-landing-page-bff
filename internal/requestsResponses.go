package internal

import "time"

type SignedURLRequest struct {
	Assets []string `json:"assets"`
}

type SignedURLResponse struct {
	URLs      map[string]string `json:"urls"`
	ExpiresIn time.Time         `json:"expires_in"`
}

type SendEmailRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}
