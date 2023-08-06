package Utils

import (
	"chatServer/apimodels"
	"chatServer/dbmodels"
	"chatServer/mongo"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDMObject(id primitive.ObjectID) (dbmodels.DirectMessageDB, error) {
	// Get the MongoDB collection
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		return dbmodels.DirectMessageDB{}, err
	}
	collection := client.Database("ProConnect").Collection("dm")

	// Create a filter for the `_id` field
	filter := bson.M{"_id": id}

	// Find the `dmobj`
	dmobj := dbmodels.DirectMessageDB{}
	err = collection.FindOne(context.TODO(), filter).Decode(&dmobj)
	if err != nil {
		return dbmodels.DirectMessageDB{}, err
	}

	// Return the `dmobj`
	return dmobj, nil
}

func GetGMObject(id primitive.ObjectID) (dbmodels.GroupMessageDB, error) {
	// Get the MongoDB collection
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		return dbmodels.GroupMessageDB{}, err
	}
	collection := client.Database("ProConnect").Collection("gm")

	// Create a filter for the `_id` field
	filter := bson.M{"_id": id}

	// Find the `gmobj`
	gmobj := dbmodels.GroupMessageDB{}
	err = collection.FindOne(context.TODO(), filter).Decode(&gmobj)
	if err != nil {
		return dbmodels.GroupMessageDB{}, err
	}

	// Return the `gmobj`
	return gmobj, nil
}

func GetUserObject(id primitive.ObjectID) (dbmodels.User, error) {
	// Get the MongoDB collection
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		return dbmodels.User{}, errors.New("Error while connecting to DB")
	}
	collection := client.Database("ProConnect").Collection("User")

	// Create a filter for the `_id` field
	filter := bson.M{"_id": id}

	// Find the `userobj`
	userobj := dbmodels.User{}
	err = collection.FindOne(context.TODO(), filter).Decode(&userobj)
	if err != nil {
		return dbmodels.User{}, errors.New("Error while getting user data")
	}

	// Return the `userobj`
	return userobj, nil
}

func GetUserGroups(userIDObjectID primitive.ObjectID) ([]primitive.ObjectID, error) {
	// Get the user object for the specified member ID
	user, err := GetUserObject(userIDObjectID)
	if err != nil {
		var emptyArr []primitive.ObjectID
		return emptyArr, err
	}
	if len(user.Groups) == 0 {
		var emptyArr []primitive.ObjectID
		return emptyArr, errors.New("User is not in any groups")
	}
	return user.Groups, nil
}

func GetGroupMessageList(groupID primitive.ObjectID) (dbmodels.GroupMessageList, error) {
	// Get the group message object for the specified group ID

	// Get the MongoDB collection
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		return dbmodels.GroupMessageList{}, errors.New("Error while connecting to DB")
	}
	collection := client.Database("ProConnect").Collection("gmlist")

	// Create a filter for the `_id` field
	filter := bson.M{"group_id": groupID}
	// Find the `grpobj`
	grpobj := dbmodels.GroupMessageList{}
	err = collection.FindOne(context.TODO(), filter).Decode(&grpobj)
	if err != nil {
		return dbmodels.GroupMessageList{}, errors.New("Error while getting grp msg data")
	}

	// Return the `grpobj`
	return grpobj, nil
}

func GetGroupObject(groupID primitive.ObjectID) (dbmodels.GroupDB, error) {
	// Get the group object for the specified group ID

	// Get the MongoDB collection
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		return dbmodels.GroupDB{}, errors.New("Error while connecting to DB")
	}
	collection := client.Database("ProConnect").Collection("groups")

	// Create a filter for the `_id` field
	filter := bson.M{"_id": groupID}

	// Find the `grpobj`
	grpobj := dbmodels.GroupDB{}

	err = collection.FindOne(context.TODO(), filter).Decode(&grpobj)
	if err != nil {
		return dbmodels.GroupDB{}, errors.New("Error while getting grp data")
	}

	// Return the `grpobj`
	return grpobj, nil
}

func GetDirectMessageList(user1, user2 primitive.ObjectID) (dbmodels.DirectMessageList, error) {
	// Get the dm chat of 2 users

	// Get the MongoDB collection
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		return dbmodels.DirectMessageList{}, errors.New("Error while connecting to DB")
	}

	collection := client.Database("ProConnect").Collection("dmlist")

	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Find the dm list record for the given user1 and user2
	filter1 := bson.M{"user1": user1, "user2": user2}
	filter2 := bson.M{"user1": user2, "user2": user1}
	var dmListObj dbmodels.DirectMessageList
	err = collection.FindOne(ctx, filter1).Decode(&dmListObj)
	if err != nil {
		if err == mongo.GetMongoErrNoDoc() {
			err = collection.FindOne(ctx, filter2).Decode(&dmListObj)
			if err != nil {
				if err == mongo.GetMongoErrNoDoc() {
					return dbmodels.DirectMessageList{}, nil
				} else {
					// There was an error finding the dm list record
					return dbmodels.DirectMessageList{}, errors.New("There was an error finding the dm list record")
				}
			}
		} else {
			// There was an error finding the dm list record
			return dbmodels.DirectMessageList{}, errors.New("There was an error finding the dm list record")
		}
	}
	// Return the `chatobj`
	return dmListObj, nil
}
func FetchAllUsers() ([]apimodels.UserAPI, error) {

	// Get the MongoDB collection
	client, err := mongo.ConnectMongoDB()
	if err != nil {
		return []apimodels.UserAPI{}, errors.New("Error while connecting to DB")
	}

	// Get the collection object
	collection := client.Database("ProConnect").Collection("User")

	// Create a cursor
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	// Iterate over the cursor and get all the users
	var users []apimodels.UserAPI
	for cursor.Next(context.TODO()) {
		var user apimodels.UserAPI
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Close the cursor
	cursor.Close(context.TODO())

	return users, nil
}
