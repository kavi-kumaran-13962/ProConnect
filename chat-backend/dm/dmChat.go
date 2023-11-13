package dm

import (
	"chatServer/Utils"
	"chatServer/apimodels"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// get dm chat is an HTTP handler to get the dm chat data from MongoDB.
func GetDMChat(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the request URL
	otherUserID := r.URL.Query().Get("otherUserID")
	if otherUserID == "" {
		http.Error(w, "Missing user IDs", http.StatusBadRequest)
		return
	}

	// Convert the user IDs to MongoDB ObjectIDs
	otherObjectID, err := primitive.ObjectIDFromHex(otherUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, _ := Utils.GetBearerToken(r)
	userId, _ := Utils.GetUserIDFromToken(token)

	userIdObjectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get the direct message list from the database
	dmList, err := Utils.GetDirectMessageList(otherObjectID, userIdObjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dmChatObj apimodels.DirectMessageChat

	dmChatObj.UserID = otherObjectID.Hex()
	userObj, err := Utils.GetUserObject(otherObjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dmChatObj.Username = userObj.Username

	for _, msgId := range dmList.Messages {
		dmObj, err := Utils.GetDMObject(msgId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		isSent := false
		if dmObj.SenderID == userIdObjectID {
			isSent = true
		}
		msgObj := apimodels.Message{
			Content:   dmObj.Content,
			Timestamp: dmObj.Timestamp,
			IsSent:    isSent,
		}
		dmChatObj.Messages = append(dmChatObj.Messages, msgObj)
	}

	// Encode the results as JSON and send them to the client
	json.NewEncoder(w).Encode(dmChatObj)
}
