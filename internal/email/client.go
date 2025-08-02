package email

import "os"

type Email struct {
	Address  string
	Password string
	SmtpHost string
	SmtpPort string
}

func NewClient() (*Email, error) {
	return &Email{
		Address:  os.Getenv("EMAIL_ADDRESS"),
		Password: os.Getenv("EMAIL_SECRET"),
		SmtpHost: os.Getenv("SMTP_HOST"),
		SmtpPort: os.Getenv("SMTP_PORT"),
	}, nil
}
