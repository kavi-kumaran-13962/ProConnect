package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserWithId struct {
	UserID    primitive.ObjectID   `bson:"_id"`
	Username  string               `bson:"username,omitempty"`
	Password  string               `bson:"password,omitempty"`
	CreatedAt time.Time            `bson:"created_at,omitempty"`
	Groups    []primitive.ObjectID `bson:"groups"`
}
