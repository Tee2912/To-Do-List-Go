package todo

import (
	"context"
	"fmt"
	"todo/db/todo/models"

	"github.com/Tee2912/To-Do-List-Go/api/proto/todolist"
	"github.com/Tee2912/To-Do-List-Go/db/todo/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "todo"

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
	findOne = func(ctx context.Context, collection *mongo.Collection, filter interface{}, p *models.ToDo) error {
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

// Get retrieves a ToDo from the database by uuid.
func Get(ctx context.Context, db *store.MongoDb, req *todolist.GetProductRequest) (*models.Product, error) {
	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
	var todo models.ToDo
	err := findOne(ctx, coll, bson.M{"uuid": req.GetUuid()}, &todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf(`todo with uuid "%s" does not exist`, req.GetUuid())
		}
		return nil, errors.Wrapf(err, `getting todo with uuid "%s"`, req.GetUuid())
	}
	return &todo, nil
}

// // Create creates a new product in the database.
// func Create(ctx context.Context, db *store.MongoDb, newProduct *models.Product) (*models.Product, error) {
// 	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
// 	newProduct.Uuid = uuidProvider()
// 	_, err := insertIntoCollection(ctx, coll, newProduct)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "inserting product")
// 	}
// 	return newProduct, nil
// }

// // Update updates a product in the database.
// func Update(ctx context.Context, db *store.MongoDb, productToUpdate *models.Product) (*models.Product, error) {
// 	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
// 	_, err := updateOne(ctx, coll, bson.M{"uuid": productToUpdate.Uuid}, bson.M{"$set": productToUpdate})
// 	if err != nil {
// 		return nil, errors.Wrapf(err, `updating product with uuid "%s"`, productToUpdate.Uuid)
// 	}
// 	return productToUpdate, nil
// }

// // Delete deletes a product from the database by uuid.
// func Delete(ctx context.Context, db *store.MongoDb, req *productcatalog.DeleteProductRequest) (*productcatalog.DeleteProductResponse, error) {
// 	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
// 	_, err := deleteOne(ctx, coll, bson.M{"uuid": req.Uuid})
// 	if err != nil {
// 		return nil, errors.Wrapf(err, `deleting product with uuid "%s"`, req.Uuid)
// 	}
// 	return &productcatalog.DeleteProductResponse{Result: "success"}, nil
// }

// // List lists all products in the database.
// func List(ctx context.Context, db *store.MongoDb, req *productcatalog.ListProductsRequest) ([]*models.Product, error) {
// 	coll := db.Client.Database(db.DatabaseName).Collection(collectionName)
// 	cur, err := find(ctx, coll, bson.M{})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "finding products")
// 	}
// 	defer cur.Close(ctx)
// 	var products []*models.Product
// 	for cur.Next(ctx) {
// 		var product models.Product
// 		if err = cur.Decode(&product); err != nil {
// 			return nil, errors.Wrap(err, "decoding product")
// 		}
// 		products = append(products, &product)
// 	}
// 	if err := cur.Err(); err != nil {
// 		return nil, errors.Wrap(err, "cursor error")
// 	}
// 	return products, nil
// }
