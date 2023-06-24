package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func SaveUser(user User) error {
	// Get the collection object
	collection := client.Database("mydb").Collection("users")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user into the collection
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// Get user details from request body
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user details
	if user.Username == "" || user.Password == "" {
		http.Error(w, "username and password are required", http.StatusBadRequest)
		return
	}

	// Save user details in database
	err = SaveUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create session for user
	// ...
}
