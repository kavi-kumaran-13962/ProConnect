package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username  string               `bson:"username,omitempty"`
	Password  string               `bson:"password,omitempty"`
	CreatedAt time.Time            `bson:"created_at,omitempty"`
	Groups    []primitive.ObjectID `bson:"groups"`
}
