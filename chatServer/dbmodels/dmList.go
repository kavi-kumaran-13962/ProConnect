package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DirectMessageList struct {
	User1       primitive.ObjectID   `bson:"user1"`
	User2       primitive.ObjectID   `bson:"user2"`
	Messages    []primitive.ObjectID `bson:"messages,omitempty"`
	LastUpdated time.Time            `bson:"last_updated"`
}
