package models

import (
	"time"
)

type GroupAPI struct {
	GroupID     string    `bson:"group_id,omitempty"`
	AdminUserID string    `bson:"admin_user_id,omitempty"`
	Members     []string  `bson:"members,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
}
