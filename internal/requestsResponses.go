/*
This file is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License.
You may obtain a copy of the license at https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt
*/

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
