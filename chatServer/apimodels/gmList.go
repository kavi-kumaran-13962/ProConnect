package apimodels

import (
	"time"
)

type GroupMessageListAPI struct {
	GroupID     string    `json:"group_id"`
	GroupName   string    `json:"groupname"`
	LastMessage string    `json:"last_message"`
	Timestamp   time.Time `json:"timestamp"`
}
