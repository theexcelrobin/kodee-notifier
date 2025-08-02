package email

import (
	"fmt"
	"net/smtp"
)

func (e *Email) Notify(email, subject, content string) error {
	msg := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, content))

	if err := smtp.SendMail(e.Host+":"+e.Port, e.Auth, e.Address, []string{email}, msg); err != nil {
		return err
	}

	return nil
}
