package email

type Email struct {
	Address  string
	Password string
	SmtpHost string
	SmtpPort string
}

func NewEmail() (*Email, error) {
	return &Email{}, nil
}
