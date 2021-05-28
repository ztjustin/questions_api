package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  int                `bson:"age"`
}

// type Users []User
