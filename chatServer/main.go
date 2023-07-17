package main

import (
	"chatServer/auth"
	"chatServer/dm"
	"chatServer/gm"
	"chatServer/websocketconnection"
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// Map to store connections for each client
var connections = make(map[string]*websocket.Conn)

// Create a new store for session management
var store = sessions.NewCookieStore([]byte("rTw3$&5z#J%G6f@Kp$7y^9rL"))

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		println("upgrade failed: ", err)
		return
	}
	// Read the client ID from the URL
	query := r.URL.Query()
	clientID := query.Get("client-id")
	println(clientID)
	if clientID == "" {
		println("client ID not found")
		conn.Close()
		return
	}
	// Add the connection to the map
	connections[clientID] = conn
	defer func() {
		// Remove the connection from the map when it is closed
		delete(connections, clientID)
		conn.Close()
	}()
	for {
		// Read a message from the client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			println("read failed:", err)
			break
		}

		// Parse the message as a JSON object
		var message map[string]string
		err = json.Unmarshal(msg, &message)
		if err != nil {
			println("parse failed:", err)
			break
		}

		// Get the recipient ID from the message
		to := message["to"]
		if to == "" {
			println("to not found")
			break
		}
		println(connections)
		// Get the recipient connection from the map
		toConn, ok := connections[to]
		if !ok {
			println("recipient not found")
			break
		}

		// Send the message to the recipient
		err = toConn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			println("write failed:", err)
			break
		}

	}
}

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
	http.Handle("/ws", auth.JWTMiddleware(http.HandlerFunc(websocketconnection.HandleConnections)))
	// signup handler
	http.HandleFunc("/signUp", auth.SignupHandler)
	// Login Handler
	http.HandleFunc("/login", auth.LoginHandler)
	// New endpoint to save direct message data
	http.HandleFunc("/saveDM", dm.SaveDMHandler)
	// New endpoint to save direct message data
	http.HandleFunc("/saveGM", gm.SaveGroupMessageHandler)
	// Call the function to create the dm collection with schema validation
	// mongo.CreateDMCollectionValidation()
	// Start the server on port 8080
	http.ListenAndServe(":8082", nil)
}
