package dm

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

// SaveDMHandler is an HTTP handler to save the direct message data to MongoDB.
func SaveDMHandler(w http.ResponseWriter, r *http.Request) {

	// Get dm details from request body
	var dm apimodels.DirectMessageAPI
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
	dmObj := dbmodels.DirectMessageDB{
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

	// Insert the msg into the collection
	result, err := collection.InsertOne(ctx, dmObj)
	if err != nil {
		http.Error(w, "cannot store dm value to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	// Get the _id of the inserted document
	id := result.InsertedID.(primitive.ObjectID)

	// Get the collection object
	collection2 := client.Database("ProConnect").Collection("dmlist")

	// Create a context with a timeout of 10 seconds
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the dm list record for the given user1 and user2
	filter1 := bson.M{"user1": senderIDObjectID, "user2": recipientIDObjectID}
	filter2 := bson.M{"user1": recipientIDObjectID, "user2": senderIDObjectID}
	var dmListObj dbmodels.DirectMessageList
	err = collection2.FindOne(ctx2, filter1).Decode(&dmListObj)
	if err != nil {
		if err == mongo.GetMongoErrNoDoc() {
			err = collection2.FindOne(ctx2, filter2).Decode(&dmListObj)
			if err != nil {
				if err == mongo.GetMongoErrNoDoc() {
					// The dm list record does not exist, so create a new one
					dmListObj = dbmodels.DirectMessageList{
						User1:       senderIDObjectID,
						User2:       recipientIDObjectID,
						Messages:    []primitive.ObjectID{},
						LastUpdated: time.Now(),
					}

					// Add the newly inserted message's _id to the dm list record
					dmListObj.Messages = append(dmListObj.Messages, id)

					// Insert the dm list record
					_, err := collection2.InsertOne(ctx2, dmListObj)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
				} else {
					// There was an error finding the dm list record
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			} else {
				// Add the newly inserted message's _id to the dm list record
				dmListObj.Messages = append(dmListObj.Messages, id)
				dmListObj.LastUpdated = time.Now()
				// The dm list record already exists, so update it
				_, err := collection2.ReplaceOne(ctx2, filter2, dmListObj)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}
		} else {
			// There was an error finding the dm list record
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// Add the newly inserted message's _id to the dm list record
		dmListObj.Messages = append(dmListObj.Messages, id)
		dmListObj.LastUpdated = time.Now()
		// The dm list record already exists, so update it
		_, err := collection2.ReplaceOne(ctx2, filter1, dmListObj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	// Return a success response
	response := map[string]string{"message": "Direct message saved successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
