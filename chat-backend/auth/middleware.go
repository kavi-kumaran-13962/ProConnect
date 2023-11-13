package auth

import (
	"chatServer/Utils"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// Middleware to check the presence of JWT token and validate that token in every request header
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the bearer token
		tokenStr, err := Utils.GetBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Validate the token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, errors.New("Invalid signing method")
			}

			// Get the signing key from a secure location
			signingKey, err := Utils.GetSigningKey()
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
