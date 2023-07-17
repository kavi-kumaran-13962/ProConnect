package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DirectMessageDB struct {
	SenderID    primitive.ObjectID `bson:"sender_id"`
	RecipientID primitive.ObjectID `bson:"recipient_id"`
	Content     string             `bson:"content"`
	Timestamp   time.Time          `bson:"timestamp"`
}
