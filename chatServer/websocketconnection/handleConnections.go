package websocketconnection

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// Map to store connections for each client
var connections = make(map[string]*websocket.Conn)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
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
