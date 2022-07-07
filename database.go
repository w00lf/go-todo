package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Database *mongo.Database
}

type Todo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func InsertedIDHEX(result *mongo.InsertOneResult) string {
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex()
	} else {
		return ""
	}
}

func (r Repository) CreateTodo(params CreateTodoParams) (Todo, error) {
	coll := r.Database.Collection("todos")
	result, err := coll.InsertOne(context.TODO(), params)
	if err != nil {
		return Todo{}, err
	}
	return Todo{Id: InsertedIDHEX(result)}, nil
}
