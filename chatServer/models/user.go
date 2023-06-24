package models

type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}
