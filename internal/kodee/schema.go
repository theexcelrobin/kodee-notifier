package main

type NotificationRequest struct {
	ID       uint    `json:"id"`
	Email    *string `json:"email"`
	Telegram *uint   `json:"telegram"`
	Whatsapp *uint   `json:"whatsapp"`
}
