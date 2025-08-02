package telegram

type Telegram struct {
	Client      string
	PhoneNumber string
}

func NewClient() (*Telegram, error) {
	return &Telegram{}, nil
}
