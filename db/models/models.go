package models

type ToDo struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Status      bool   `bson:"status"`
	Id          string `bson:"id"`
}

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}
