package gm

import (
	"chatServer/apimodels"
	"chatServer/dbmodels"

	"chatServer/mongo"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveGroupMessageHandler is an HTTP handler to save the group message data to MongoDB.
func SaveGroupMessageHandler(w http.ResponseWriter, r *http.Request) {

	// Get group message details from request body
	var groupMessage apimodels.GroupMessageAPI
	err := json.NewDecoder(r.Body).Decode(&groupMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate required fields
	if groupMessage.SenderID == "" || groupMessage.GroupID == "" || groupMessage.Content == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Set the timestamp to the current time
	timestamp := time.Now()
	timestampString := timestamp.Format("2006-01-02T15:04:05Z")
	timestamp, err = time.Parse("2006-01-02T15:04:05Z", timestampString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	groupIDObjectID, err := primitive.ObjectIDFromHex(groupMessage.GroupID)
	senderIDObjectID, err := primitive.ObjectIDFromHex(groupMessage.SenderID)

	// Create a new GroupMessage struct
	groupMessageObj := dbmodels.GroupMessageDB{
		GroupID:   groupIDObjectID,
		SenderID:  senderIDObjectID,
		Content:   groupMessage.Content,
		Timestamp: timestamp,
	}

	// Call the ConnectMongoDB function
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		http.Error(w, "cannot connect to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	// Get the collection object
	collection := client.Database("ProConnect").Collection("gm")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the msg into the collection
	result, err := collection.InsertOne(ctx, groupMessageObj)
	if err != nil {
		http.Error(w, "cannot store group message value to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Get the _id of the inserted document
	id := result.InsertedID.(primitive.ObjectID)

	// Get the collection object
	collection2 := client.Database("ProConnect").Collection("gmlist")

	// Create a context with a timeout of 10 seconds
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the gm list record for the given groupId
	filter := bson.M{"group_id": groupIDObjectID}
	var gmListObj dbmodels.GroupMessageList
	err = collection2.FindOne(ctx2, filter).Decode(&gmListObj)
	if err != nil {
		if err == mongo.GetMongoErrNoDoc() {
			// The gm list record does not exist, so create a new one
			gmListObj = dbmodels.GroupMessageList{
				GroupID:     groupIDObjectID,
				Messages:    []primitive.ObjectID{},
				LastUpdated: time.Now(),
			}

			// Add the newly inserted message's _id to the gm list record
			gmListObj.Messages = append(gmListObj.Messages, id)

			// Insert the gm list record
			_, err := collection2.InsertOne(ctx2, gmListObj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

		} else {
			// There was an error finding the gm list record
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// Add the newly inserted message's _id to the gm list record
		gmListObj.Messages = append(gmListObj.Messages, id)
		gmListObj.LastUpdated = time.Now()
		// The gm list record already exists, so update it
		_, err := collection2.ReplaceOne(ctx2, filter, gmListObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Return a success response
	response := map[string]string{"message": "Group message saved successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
