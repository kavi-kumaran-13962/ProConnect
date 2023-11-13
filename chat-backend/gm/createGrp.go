package gm

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
)

// CreateGroupHandler is an HTTP handler to create a new group.
func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {

	// Get group details from request body
	var group apimodels.GroupAPI
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if group.AdminUserID == "" || len(group.Members) == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Set the timestamp to the current time
	timestamp := time.Now()

	adminUserId, err := primitive.ObjectIDFromHex(group.AdminUserID)

	// Create a new Group struct
	groupObj := dbmodels.GroupDB{
		GroupName:   group.GroupName,
		AdminUserID: adminUserId,
		Members:     make([]primitive.ObjectID, len(group.Members)),
		CreatedAt:   timestamp,
	}

	// Iterate over the members and convert them to ObjectIDs
	for i, memberID := range group.Members {
		memberIDObjectID, err := primitive.ObjectIDFromHex(memberID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		groupObj.Members[i] = memberIDObjectID
	}

	// Call the ConnectMongoDB function
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		http.Error(w, "cannot connect to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	// Get the collection object
	collection := client.Database("ProConnect").Collection("groups")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the group into the collection
	result, err := collection.InsertOne(ctx, groupObj)
	// Get the _id of the inserted document
	groupId := result.InsertedID.(primitive.ObjectID)

	// Iterate over the members and convert them to ObjectIDs
	for _, memberID := range group.Members {
		memberIDObjectID, err := primitive.ObjectIDFromHex(memberID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get the user object for the specified member ID
		user, err := Utils.GetUserObject(memberIDObjectID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the group ID to the user's `Groups` array
		user.Groups = append(user.Groups, groupId)

		// Create a context with a timeout of 10 seconds
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		collection = client.Database("ProConnect").Collection("User")

		filter := bson.M{"_id": memberIDObjectID}

		_, err = collection.ReplaceOne(ctx, filter, user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	if err != nil {
		http.Error(w, "cannot store group value to mongoDB", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Return a success response
	response := map[string]string{"message": "Group created successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getAllUsers is an HTTP handler to get all users for adding users to group.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersList, err := Utils.FetchAllUsers()
	if err != nil {
		http.Error(w, "cannot get users list", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	// Encode the results as JSON and send them to the client
	json.NewEncoder(w).Encode(usersList)
}
