package apimodels

import (
	"time"
)

type GroupMessageAPI struct {
	GroupID   string    `json:"group_id"`
	SenderID  string    `json:"sender_id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
