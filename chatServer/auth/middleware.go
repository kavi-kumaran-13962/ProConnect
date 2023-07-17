package auth

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// Middleware to check the presence of JWT token and validate that token in every request header
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			// Token is not present in the header
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token from the header
		tokenStr := authHeader[len("Bearer "):]

		// Validate the token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, errors.New("Invalid signing method")
			}

			// Get the signing key from a secure location
			signingKey, err := getSigningKey()
			if err != nil {
				return nil, err
			}

			// Return the signing key
			return signingKey, nil
		})

		// If the token is invalid
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}

func getSigningKey() ([]byte, error) {
	// Get the signing key from a secure location
	// ...

	return signingKey, nil
}
