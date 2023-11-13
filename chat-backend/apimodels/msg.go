package apimodels

import "time"

type Message struct {
	Content   string    `json:"content"`
	IsSent    bool      `json:"isSent"`
	Timestamp time.Time `json:"timestamp"`
}
