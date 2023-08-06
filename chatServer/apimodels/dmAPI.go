package apimodels

import (
	"time"
)

type DirectMessageAPI struct {
	SenderID    string    `json:"sender_id"`
	RecipientID string    `json:"recipient_id"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
}
