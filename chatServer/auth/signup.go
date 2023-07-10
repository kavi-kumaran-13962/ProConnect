package auth

import (
	"chatServer/models"
	"chatServer/mongo"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func saveUser(user models.User) error {

	// Call the ConnectMongoDB function
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Get the collection object
	collection := client.Database("ProConnect").Collection("User")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user into the collection
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	// Get user details from request body
	var user models.User
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

	// Get the current time
	now := time.Now()

	// add createdAt
	user.CreatedAt = now

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the hashed password
	user.Password = string(hashedPassword)

	// Save user details in database
	err = saveUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Successfully created user
	w.WriteHeader(http.StatusCreated)
}
