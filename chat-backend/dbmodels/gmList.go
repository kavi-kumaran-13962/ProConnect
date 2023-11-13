package dbmodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupMessageList struct {
	GroupID     primitive.ObjectID   `bson:"group_id"`
	Messages    []primitive.ObjectID `bson:"messages,omitempty"`
	LastUpdated time.Time            `bson:"last_updated"`
}
