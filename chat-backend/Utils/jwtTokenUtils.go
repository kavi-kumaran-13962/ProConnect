package Utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// GetBearerToken extracts the bearer token from the Authorization header
func GetBearerToken(r *http.Request) (string, error) {
	// Get the token from the Authorization header
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		// Token is not present in the header
		return "", errors.New("Unauthorized")
	}

	// Extract the token from the header
	tokenStr := authHeader[len("Bearer "):]

	return tokenStr, nil
}

func GetUserIDFromToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("Invalid signing method")
		}

		// Get the signing key
		signingKey, _ := GetSigningKey()

		return signingKey, nil
	})

	// Check if the token is valid
	if err != nil {
		return "", err
	}

	// Get the user ID from the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	// Return the user ID
	userId, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("invalid user ID")
	}

	return userId, nil
}

func GetUserNameFromToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("Invalid signing method")
		}

		// Get the signing key
		signingKey, _ := GetSigningKey()

		return signingKey, nil
	})

	// Check if the token is valid
	if err != nil {
		return "", err
	}

	// Get the username from the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	// Return the user ID
	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("invalid user username")
	}

	return username, nil
}

func GetSigningKey() ([]byte, error) {

	// JWT signing key
	var signingKey = []byte("your-secret-key")

	// Get the signing key from a secure location
	// ...

	return signingKey, nil
}
