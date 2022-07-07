package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppConfiguration struct {
	MongoDBURI          string
	MongoDBDatabaseName string
}

func main() {
	var config AppConfiguration
	err := envconfig.Process("", &config)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	database := client.Database(config.MongoDBDatabaseName)
	repo := Repository{Database: database}
	r := mux.NewRouter()
	r.HandleFunc("/todos", repo.CreateTodoHandler).
		Methods("POST")
	// r.HandleFunc("/todos/{id:[0-9]+}/", ShowTodoHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
