package home

import (
	"chatServer/Utils"
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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetDMListHandler is an HTTP handler to get the direct message list for a given user ID.
func GetDMListHandler(w http.ResponseWriter, r *http.Request) {

	// Get the user ID from the request URL
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	// Convert the user ID to a MongoDB ObjectID
	userIDObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a filter for the `dmlist` collection
	filter := bson.M{"$or": []bson.M{
		{"user1": userIDObjectID},
		{"user2": userIDObjectID},
	}}

	// Create a sort struct
	sort := options.FindOptions{
		Sort: bson.M{"last_updated": -1},
	}

	// Create a pointer to the sort struct
	sortPtr := &sort

	// Get the `dmlist` collection
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		http.Error(w, "cannot connect to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	collection := client.Database("ProConnect").Collection("dmlist")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all the matching documents
	cursor, err := collection.Find(ctx, filter, sortPtr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Iterate over the results and decode them
	var dmLists []apimodels.DirectMessageListAPI
	for cursor.Next(ctx) {
		var dmList dbmodels.DirectMessageList
		err := cursor.Decode(&dmList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Get the last message ID from the `messages` array
		lastMessageID := dmList.Messages[len(dmList.Messages)-1]
		msgObj, err := Utils.GetDMObject(lastMessageID)
		var userId primitive.ObjectID
		if dmList.User1 == userIDObjectID {
			userId = dmList.User2
		} else {
			userId = dmList.User1
		}
		userObj, err := Utils.GetUserObject(userId)
		dmListAPI := apimodels.DirectMessageListAPI{
			UserID:      userId.Hex(),
			UserName:    userObj.Username,
			LastMessage: msgObj.Content,
			Timestamp:   time.Now(),
		}
		dmLists = append(dmLists, dmListAPI)
	}
	// Close the cursor
	cursor.Close(ctx)

	// Encode the results as JSON and send them to the client
	json.NewEncoder(w).Encode(dmLists)
}
