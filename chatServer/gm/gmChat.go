package gm

import (
	"chatServer/Utils"
	"chatServer/apimodels"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// get Gm chat is an HTTP handler to get the Gm chat data from MongoDB.
func GetGMChat(w http.ResponseWriter, r *http.Request) {
	// Get the grp ID from the request URL
	groupId := r.URL.Query().Get("groupId")
	if groupId == "" {
		http.Error(w, "Missing group IDs", http.StatusBadRequest)
		return
	}

	// Convert the user IDs to MongoDB ObjectIDs
	grpObjectID, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the group message list from the database
	gmList, err := Utils.GetGroupMessageList(grpObjectID)
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

	var gmChatObj apimodels.GroupMessageChat
	gmChatObj.GroupID = groupId
	grpObj, _ := Utils.GetGroupObject(grpObjectID)
	gmChatObj.Groupname = grpObj.GroupName
	for _, msgId := range gmList.Messages {
		gmObj, err := Utils.GetGMObject(msgId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		isSent := false
		if gmObj.SenderID == userIdObjectID {
			isSent = true
		}
		msgObj := apimodels.Message{
			Content:   gmObj.Content,
			Timestamp: gmObj.Timestamp,
			IsSent:    isSent,
		}
		gmChatObj.Messages = append(gmChatObj.Messages, msgObj)
	}

	// Encode the results as JSON and send them to the client
	json.NewEncoder(w).Encode(gmChatObj)
}
