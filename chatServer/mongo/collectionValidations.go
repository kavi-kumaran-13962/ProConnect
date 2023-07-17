package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDMCollectionValidation() {
	// Call the ConnectMongoDB function
	client, err := ConnectMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	// Define the JSON schema for dm collection
	var dmCollectionSchema = bson.M{
		"bsonType": "object",
		"required": []string{"sender_id", "recipient_id", "content", "timestamp"},
		"properties": bson.M{
			"sender_id": bson.M{
				"bsonType":    "objectId",
				"description": "Sender ID is required and must be a objectId",
			},
			"recipient_id": bson.M{
				"bsonType":    "objectId",
				"description": "Recipient ID is required and must be a objectId",
			},
			"content": bson.M{
				"bsonType":    "string",
				"description": "Content is required and must be a string",
			},
			"timestamp": bson.M{
				"bsonType":    "date",
				"description": "Timestamp is required and must be a date",
			},
		},
	}

	// Create the collection with the JSON schema validation
	collectionOptions := options.CreateCollection().SetValidator(dmCollectionSchema)
	if err = client.Database("ProConnect").CreateCollection(context.Background(), "dm", collectionOptions); err != nil {
		// Check if the collection already exists
		if mErr, ok := err.(mongo.CommandError); ok && mErr.Code == 48 {
			fmt.Println("Collection dm already exists.")
		} else {
			fmt.Println(err)
		}
	}
}
