package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Status      bool   `bson:"status"`
	Id          string `bson:"id"`
}

type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
}
