package whatsapp

import (
	"context"
	"fmt"
	"os"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"

	_ "github.com/mattn/go-sqlite3"
	"github.com/skip2/go-qrcode"
)

type Whatsapp struct {
	Client *whatsmeow.Client
}

func NewClient() (*Whatsapp, error) {
	ctx := context.Background()

	container, err := sqlstore.New(ctx, "sqlite3", fmt.Sprintf("file:%s.db?_foreign_keys=on", os.Getenv("WHATSAPP_SESSION")), nil)
	if err != nil {
		return nil, fmt.Errorf("database error: %s", err.Error())
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		return nil, fmt.Errorf("device store error: %s", err.Error())
	}

	client := whatsmeow.NewClient(deviceStore, nil)
	if client.Store.ID == nil {
		qrChan, err := client.GetQRChannel(context.Background())
		if err != nil {
			return nil, fmt.Errorf("qr generation error: %s", err.Error())
		}

		if err := client.Connect(); err != nil {
			return nil, fmt.Errorf("client connection error: %s", err.Error())
		}

		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Scan this QR code with your WhatsApp mobile app:")

				qr, err := qrcode.New(evt.Code, qrcode.Medium)
				if err == nil {
					fmt.Println(qr.ToSmallString(false))
				}

				if err := qrcode.WriteFile(evt.Code, qrcode.Low, 16, "qr.png"); err == nil {
					fmt.Println("QR code also saved as qr.png")
				}
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		if err = client.Connect(); err != nil {
			return nil, fmt.Errorf("client connection error: %s", err.Error())
		}
	}

	return &Whatsapp{
		Client: client,
	}, nil
}
