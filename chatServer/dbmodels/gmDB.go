package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupMessageDB struct {
	GroupID   primitive.ObjectID `bson:"group_id,omitempty"`
	SenderID  primitive.ObjectID `bson:"sender_id,omitempty"`
	Content   string             `bson:"content,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
}
