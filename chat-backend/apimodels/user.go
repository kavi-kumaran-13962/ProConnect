package apimodels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAPI struct {
	UserID   primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username,omitempty"`
}
