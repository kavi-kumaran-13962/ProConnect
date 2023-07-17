package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupDB struct {
	GroupID     primitive.ObjectID   `bson:"group_id,omitempty"`
	AdminUserID primitive.ObjectID   `bson:"admin_user_id,omitempty"`
	Members     []primitive.ObjectID `bson:"members,omitempty"`
	CreatedAt   time.Time            `bson:"created_at,omitempty"`
}
