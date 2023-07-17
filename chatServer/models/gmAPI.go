package models

import (
	"time"
)

type GroupMessageAPI struct {
	GroupID   string    `json:"group_id,omitempty"`
	SenderID  string    `json:"sender_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
