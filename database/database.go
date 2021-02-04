package database

import (
	"context"
	"go-graphql-test/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB Type
type DB struct {
	client *mongo.Client
}

// Db is the exported variable that allows
// other internal packages to access Mongo client
var Db *DB

// Connect Fn
func Connect(uri string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Print("MongoDB Successfully Connected!")
	Db = &DB{
		client: client,
	}
}

// CreateToDo Fn
func (db *DB) CreateToDo(input *model.NewTodo) (*model.Todo, error) {
	collection := db.client.Database("test").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	newTodo := &model.Todo{
		ID:     res.InsertedID.(primitive.ObjectID).Hex(),
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}
	return newTodo, nil
}

// FindTodos Fn
func (db *DB) FindTodos() ([]*model.Todo, error) {
	collection := db.client.Database("test").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var todos []*model.Todo
	for cur.Next(ctx) {
		var todo *model.Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return todos, nil
}
