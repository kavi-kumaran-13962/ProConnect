package dm

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

// SaveDMHandler is an HTTP handler to save the direct message data to MongoDB.
func SaveDMHandler(w http.ResponseWriter, r *http.Request) {

	// Get dm details from request body
	var dm models.DirectMessageAPI
	err := json.NewDecoder(r.Body).Decode(&dm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if dm.SenderID == "" || dm.RecipientID == "" || dm.Content == "" {
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

	recipientIDObjectID, err := primitive.ObjectIDFromHex(dm.RecipientID)
	senderIDObjectID, err := primitive.ObjectIDFromHex(dm.SenderID)

	// Create a new DirectMessage struct
	dmObj := models.DirectMessageDB{
		SenderID:    senderIDObjectID,
		RecipientID: recipientIDObjectID,
		Content:     dm.Content,
		Timestamp:   timestamp,
	}

	// Call the ConnectMongoDB function
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		http.Error(w, "cannot connect to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	// Get the collection object
	collection := client.Database("ProConnect").Collection("dm")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user into the collection
	_, err = collection.InsertOne(ctx, dmObj)
	if err != nil {
		http.Error(w, "cannot store dm value to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Return a success response
	response := map[string]string{"message": "Direct message saved successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
