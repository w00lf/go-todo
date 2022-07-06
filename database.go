package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Client *mongo.Client
}

type Todo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func (r Repository) CreateTodo(params CreateTodoParams) (Todo, error) {
	coll := r.Client.Database("sample_mflix").Collection("todos")
	result, err := coll.InsertOne(context.TODO(), params)
	if err != nil {
		return Todo{}, err
	}
	return Todo{Id: fmt.Sprintf("%v", result.InsertedID)}, nil
}
