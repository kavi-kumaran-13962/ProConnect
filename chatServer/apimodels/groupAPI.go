package apimodels

import (
	"time"
)

type GroupAPI struct {
	GroupName   string    `json:"groupname,omitempty"`
	AdminUserID string    `json:"admin_user_id,omitepty"`
	Members     []string  `json:"members,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
