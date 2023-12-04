package user

import (
	"context"
	"fmt"

	"github.com/Tee2912/To-Do-List-Go/api/proto/auth"
	"github.com/Tee2912/To-Do-List-Go/db"
	"github.com/Tee2912/To-Do-List-Go/db/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const collectionName = "user"

// Cursor is an interface that defines the methods necessary for iterating
// over query results in a data layer.
// This interface is particularly useful for simplifying unit tests
// by allowing the implementation of mock cursors that can be used
// for testing data retrieval and manipulation operations.
type Cursor interface {
	Decode(interface{}) error
	Err() error
	Close(context.Context) error
	Next(context.Context) bool
}

type cursorWrapper struct {
	*mongo.Cursor
}

// For ease of unit testing.
var (
	uuidProvider         = uuid.NewString
	insertIntoCollection = func(ctx context.Context, collection *mongo.Collection, document interface{}) (*mongo.InsertOneResult, error) {
		return collection.InsertOne(ctx, document)
	}
	find = func(ctx context.Context, collection *mongo.Collection, filter interface{}) (Cursor, error) {
		cur, err := collection.Find(ctx, filter)
		return &cursorWrapper{cur}, err
	}
	findOne = func(ctx context.Context, collection *mongo.Collection, filter interface{}, p *models.User) error {
		sr := collection.FindOne(ctx, filter)
		return sr.Decode(p)
	}
	updateOne = func(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
		return collection.UpdateOne(ctx, filter, update)
	}
	deleteOne = func(ctx context.Context, collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
		return collection.DeleteOne(ctx, filter)
	}
)

// Sign up a new user
func SignUpUser(ctx context.Context, db *db.MongoDb, newUser *models.User) (*auth.RegisterResponse, error) {
	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)

	_, err := insertIntoCollection(ctx, coll, newUser)
	if err != nil {
		return nil, errors.Wrap(err, "inserting user")
	}
	return &auth.RegisterResponse{Message: "success"}, nil
}

// Read user
func Get(ctx context.Context, db *db.MongoDb, req *auth.LoginRequest) (*models.User, error) {
	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
	var user models.User
	err := findOne(ctx, coll, bson.M{"username": req.GetUsername()}, &user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(`todo with id "%s" does not exist`, req.GetUsername())
		}
		return nil, errors.Wrapf(err, `getting todo with id "%s"`, req.GetUsername())
	}
	return &user, nil
}
