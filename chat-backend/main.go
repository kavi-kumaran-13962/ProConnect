package main

import (
	"chatServer/auth"
	"chatServer/dm"
	"chatServer/gm"
	"chatServer/home"
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
	http.Handle("/ws", auth.JWTMiddleware(http.HandlerFunc(websocketconnection.HandleConnections)))
	// signup handler
	http.HandleFunc("/signUp", auth.SignupHandler)
	// Login Handler
	http.HandleFunc("/login", auth.LoginHandler)
	// endpoint to save direct message data
	http.Handle("/sendDM", auth.JWTMiddleware(http.HandlerFunc(dm.SaveDMHandler)))
	// endpoint to grp message data
	http.Handle("/sendGM", auth.JWTMiddleware(http.HandlerFunc(gm.SaveGroupMessageHandler)))
	// endpoint to get all users
	http.Handle("/users", auth.JWTMiddleware(http.HandlerFunc(gm.GetAllUsers)))
	// endpoint to create grp
	http.Handle("/createGrp", auth.JWTMiddleware(http.HandlerFunc(gm.CreateGroupHandler)))
	// endpoint to get DM List
	http.Handle("/getDMList", auth.JWTMiddleware(http.HandlerFunc(home.GetDMListHandler)))
	// endpoint to get GM List
	http.Handle("/getGMList", auth.JWTMiddleware(http.HandlerFunc(home.GetGroupMessageListHandler)))
	// endpoint to get group chat
	http.Handle("/getGMChat", auth.JWTMiddleware(http.HandlerFunc(gm.GetGMChat)))
	// endpoint to get dm chat
	http.Handle("/getDMChat", auth.JWTMiddleware(http.HandlerFunc(dm.GetDMChat)))

	// Call the function to create the dm collection with schema validation
	// mongo.CreateDMCollectionValidation()
	// Start the server on port 8080
	http.ListenAndServe(":8082", nil)
}
