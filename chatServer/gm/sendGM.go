package gm

import (
	"chatServer/models"
	"chatServer/mongo"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SaveGroupMessageHandler is an HTTP handler to save the group message data to MongoDB.
func SaveGroupMessageHandler(w http.ResponseWriter, r *http.Request) {

	// Get group message details from request body
	var groupMessage models.GroupMessageAPI
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
	groupMessageObj := models.GroupMessageDB{
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

	// Insert the user into the collection
	_, err = collection.InsertOne(ctx, groupMessageObj)
	if err != nil {
		http.Error(w, "cannot store group message value to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Return a success response
	response := map[string]string{"message": "Group message saved successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
