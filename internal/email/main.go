package email

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Email struct {
	Address  string
	Password string
	SmtpHost string
	SmtpPort string
}

func NewEmail() (*Email, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err.Error())
	}

	return &Email{}, nil
}
