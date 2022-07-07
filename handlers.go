package main

import (
	"encoding/json"
	"net/http"
)

type CreateTodoParams struct {
	Name        string `json:"name", required:"true"`
	Description string `json:"description", required:"true"`
	Priority    int    `json:"priority", required:"true"`
}

func respondWithError(w http.ResponseWriter, status int, err error) {
	respondWithJSON(w, status, map[string]string{"error": err.Error()})
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func (repo *Repository) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todoParams CreateTodoParams
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todoParams); err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	todo, err := repo.CreateTodo(todoParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"id": todo.Id})
}
