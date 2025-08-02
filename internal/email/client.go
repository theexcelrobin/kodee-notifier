package email

import (
	"net/smtp"
	"os"
)

type Email struct {
	Address  string
	Password string
	Host     string
	Port     string
	Auth     smtp.Auth
}

func NewClient() (*Email, error) {
	e := &Email{
		Address:  os.Getenv("EMAIL_ADDRESS"),
		Password: os.Getenv("EMAIL_SECRET"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
	}

	e.Auth = smtp.PlainAuth("", e.Address, e.Password, e.Host)
	return e, nil
}
