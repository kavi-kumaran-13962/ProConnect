package home

import (
	"chatServer/Utils"
	"chatServer/apimodels"
	"sort"

	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetGroupMessageListHandler is an HTTP handler to get the group message list for a given user ID.
func GetGroupMessageListHandler(w http.ResponseWriter, r *http.Request) {

	// Get the user ID from the request URL
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	// Convert the user ID to a MongoDB ObjectID
	userIDObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the list of groups that the user is part of
	groups, err := Utils.GetUserGroups(userIDObjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a slice to store the group message lists
	groupMessageLists := []apimodels.GroupMessageListAPI{}

	// Iterate over the groups and fetch the group message lists
	for _, groupID := range groups {
		// Get the group message list from the database
		groupMessageList, err := Utils.GetGroupMessageList(groupID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get the last message ID from the `messages` array
		lastMessageID := groupMessageList.Messages[len(groupMessageList.Messages)-1]
		msgObj, err := Utils.GetGMObject(lastMessageID)
		groupObj, err := Utils.GetGroupObject(groupID)
		groupMessageListAPI := apimodels.GroupMessageListAPI{
			GroupID:     groupMessageList.GroupID.Hex(),
			GroupName:   groupObj.GroupName,
			LastMessage: msgObj.Content,
			Timestamp:   groupMessageList.LastUpdated,
		}
		groupMessageLists = append(groupMessageLists, groupMessageListAPI)
	}

	// Sort the group message lists by `Timestamp`
	sort.Slice(groupMessageLists, func(i, j int) bool {
		return groupMessageLists[i].Timestamp.Before(groupMessageLists[j].Timestamp)
	})

	// Encode the results as JSON and send them to the client
	json.NewEncoder(w).Encode(groupMessageLists)
}
