package main

import (
	"chatServer/auth"
	"chatServer/websocketconnection"
	"net/http"
)

func main() {
	// serves client html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "client.html")
		}
	})
	// serves client js
	http.HandleFunc("/client.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client.js")
	})
	// serves client css file
	http.HandleFunc("/client.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client.css")
	})
	// Handle WebSocket connections at the "/ws" endpoint
	http.HandleFunc("/ws", websocketconnection.HandleConnections)
	// signup handler
	http.HandleFunc("/signUp", auth.SignupHandler)
	// Login Handler
	http.HandleFunc("/login", auth.LoginHandler)
	// Start the server on port 8080
	http.ListenAndServe(":8082", nil)
}
