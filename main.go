package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", CreateTodoHandler).
		Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}/", ShowTodoHandler).Methods("GET")

	mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
