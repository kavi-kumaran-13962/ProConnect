package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupDB struct {
	GroupName   string               `bson:"groupname,omitempty"`
	AdminUserID primitive.ObjectID   `bson:"admin_user_id,omitempty"`
	Members     []primitive.ObjectID `bson:"members,omitempty"`
	CreatedAt   time.Time            `bson:"created_at,omitempty"`
}
