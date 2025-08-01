package whatsapp

import (
	"context"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func (w *Whatsapp) notify(phone, text string) error {
	recipientJID := types.NewJID(phone, types.DefaultUserServer)

	message := &waE2E.Message{
		Conversation: proto.String(text),
	}

	_, err := w.Client.SendMessage(context.Background(), recipientJID, message)
	return err
}
