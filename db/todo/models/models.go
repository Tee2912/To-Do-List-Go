package models

type ToDo struct {
	Id          string `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Status      bool   `bson:"status"`
}
