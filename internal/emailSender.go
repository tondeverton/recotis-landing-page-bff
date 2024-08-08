/*
This file is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License.
You may obtain a copy of the license at https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt
*/

package internal

import (
	"net/smtp"
)

func SendEmail(subject, body string) error {
	smtpUsername := GetSmtpUsername()
	smtpPassword := GetSmtpPassword()
	smtpServer := GetSmtpServer()
	smtpServerPort := GetSmtpServerPort()
	emailTarget := GetEmailTarget()

	auth := smtp.PlainAuth(
		"",
		smtpUsername,
		smtpPassword,
		smtpServer,
	)

	msg := "To: " + emailTarget + "\r\n" +
		"Subject: [Landing Page] " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n"

	err := smtp.SendMail(
		smtpServer+":"+smtpServerPort,
		auth,
		smtpUsername,
		[]string{emailTarget},
		[]byte(msg),
	)

	return err
}
