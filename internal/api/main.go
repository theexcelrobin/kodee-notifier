package api

import (
	"os"

	"github.com/theexcelrobin/kodee-notifier/internal/email"
	"github.com/theexcelrobin/kodee-notifier/internal/telegram"
	"github.com/theexcelrobin/kodee-notifier/internal/whatsapp"
)

type Api struct {
	Address string
	Port    string
}

func NewApi(e *email.Email, t *telegram.Telegram, w *whatsapp.Whatsapp) (*Api, error) {
	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	return &Api{
		Address: address,
		Port:    port,
	}, nil
}

func (a *Api) Spawn() {
	//
}
