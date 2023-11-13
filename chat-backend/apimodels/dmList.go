package apimodels

import (
	"time"
)

type DirectMessageListAPI struct {
	UserID      string    `json:"user_id"`
	UserName    string    `json:"username"`
	LastMessage string    `json:"last_message"`
	Timestamp   time.Time `json:"timestamp"`
}
