package main

import (
	"github.com/joho/godotenv"
	"github.com/theexcelrobin/kodee-notifier/internal/api"
	"github.com/theexcelrobin/kodee-notifier/internal/email"
	"github.com/theexcelrobin/kodee-notifier/internal/logger"
	"github.com/theexcelrobin/kodee-notifier/internal/telegram"
	"github.com/theexcelrobin/kodee-notifier/internal/whatsapp"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	l, err := logger.NewLogger()
	if err != nil {
		panic(err)
	} else {
		defer l.LogFile.Close()
	}

	e, err := email.NewClient()
	if err != nil {
		panic(err)
	}

	t, err := telegram.NewClient()
	if err != nil {
		panic(err)
	}

	w, err := whatsapp.NewClient()
	if err != nil {
		panic(err)
	} else {
		defer w.Client.Disconnect()
	}

	a, err := api.NewApi(e, t, w)
	if err != nil {
		panic(err)
	}

	a.SpawnServer()
}
