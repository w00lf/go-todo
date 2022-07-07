package main_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/w00lf/go-todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func executeRequest(req *http.Request, router *mux.Router, repo Repository) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

var _ = Describe("Handlers", func() {
	var repo Repository
	var router *mux.Router
	var todoColl *mongo.Collection

	BeforeEach(func() {
		var config AppConfiguration
		err := envconfig.Process("", &config)
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBURI))
		if err != nil {
			panic(err)
		}
		database := client.Database(config.MongoDBDatabaseName)
		repo = Repository{Database: database}
		router = mux.NewRouter()
		todoColl = repo.Database.Collection("todos")
		// runtime.Breakpoint()
		_, err = todoColl.DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			panic(err)
		}
	})
	Describe("create todo", func() {
		Context("with valid params passed", func() {
			// https://onsi.github.io/ginkgo/#testing-an-in-process-asynchronous-service - async assertions
			It("should create a todo entry", func() {
				todoCount := func() int {
					count, err := todoColl.EstimatedDocumentCount(context.TODO())
					if err != nil {
						panic(err)
					}
					return int(count)
				}
				Expect(todoCount()).To(Equal(0))
				var jsonStr = []byte(`{"name":"test product", "description": "Lorem ipsum test test", "priority": 1}`)
				req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonStr))
				router.HandleFunc("/todos", repo.CreateTodoHandler)
				response := executeRequest(req, router, repo)
				Expect(todoCount()).To(Equal(1))
				Expect(response.Body.String()).To(ContainSubstring(`{"id":"`))
			})
		})
	})
})
