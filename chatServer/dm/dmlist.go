package dm

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserID   string `bson:"user_id,omitempty"`
	Username string `bson:"username,omitempty"`
}

type DirectMessage struct {
	MessageID   string    `bson:"message_id,omitempty"`
	SenderID    string    `bson:"sender_id,omitempty"`
	RecipientID string    `bson:"recipient_id,omitempty"`
	Content     string    `bson:"content,omitempty"`
	Timestamp   time.Time `bson:"timestamp,omitempty"`
}

type LastMessage struct {
	Content   string    `json:"content,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []*UserData `json:"data,omitempty"`
}

type UserData struct {
	UserID      string       `json:"user_id,omitempty"`
	Username    string       `json:"username,omitempty"`
	LastMessage *LastMessage `json:"last_message,omitempty"`
}

func GetChatsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve chats from MongoDB and construct the response
	chats, err := RetrieveChatsFromDB("user1") // Replace "user1" with the desired user ID
	if err != nil {
		sendErrorResponse(w, "Failed to retrieve chats")
		return
	}

	response := APIResponse{
		Status:  "success",
		Message: "Chats retrieved successfully",
		Data:    chats,
	}

	sendJSONResponse(w, http.StatusOK, response)
}

func RetrieveChatsFromDB(userID string) ([]*UserData, error) {
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	// Get the MongoDB collection for users
	userCollection := client.Database("your_database_name").Collection("users")

	// Get the MongoDB collection for direct messages
	dmCollection := client.Database("your_database_name").Collection("direct_messages")

	// Query MongoDB to retrieve the user's direct messages
	cursor, err := dmCollection.Find(ctx, bson.M{"sender_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the retrieved direct messages and collect relevant data
	userData := []*UserData{}
	for cursor.Next(ctx) {
		var dm DirectMessage
		err := cursor.Decode(&dm)
		if err != nil {
			return nil, err
		}

		// Get the recipient's username
		var recipient User
		err = userCollection.FindOne(ctx, bson.M{"user_id": dm.RecipientID}).Decode(&recipient)
		if err != nil {
			return nil, err
		}

		// Create UserData with necessary information
		userData = append(userData, &UserData{
			UserID:   dm.RecipientID,
			Username: recipient.Username,
			LastMessage: &LastMessage{
				Content:   dm.Content,
				Timestamp: dm.Timestamp,
			},
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return userData, nil
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	response := APIResponse{
		Status:  "error",
		Message: message,
	}
	sendJSONResponse(w, http.StatusInternalServerError, response)
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Failed to send JSON response: %v", err)
	}
}
