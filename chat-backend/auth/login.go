package auth

import (
	"chatServer/Utils"
	"chatServer/dbmodels"
	"chatServer/mongo"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func loginUser(username, password string) (string, error) {
	// Call the ConnectMongoDB function
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Get the collection object
	collection := client.Database("ProConnect").Collection("User")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve the user from the collection
	var user dbmodels.UserWithId
	cursor, err := collection.Find(ctx, bson.M{"username": username})
	if err != nil {
		return "", err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("invalid username")
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("wrong password")
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
		"userId":   user.UserID,                           // Add the user ID to the token
	})
	signingKey, _ := Utils.GetSigningKey()
	// Sign the token with the secret key
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return tokenString, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get user credentials from request body
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user credentials
	if credentials.Username == "" || credentials.Password == "" {
		http.Error(w, "username and password are required", http.StatusBadRequest)
		return
	}

	// Login the user and get the JWT token
	token, err := loginUser(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Send the token as the response
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
