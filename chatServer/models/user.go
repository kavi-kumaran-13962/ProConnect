package models

import (
	"time"
)

type User struct {
	Username  string    `bson:"username,omitempty"`
	Password  string    `bson:"password,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
}
