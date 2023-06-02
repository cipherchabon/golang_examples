package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cypherchabon/golang_examples/todo_server/models"
	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := make([]models.Todo, len(models.Store))

	i := 0
	for _, todo := range models.Store {
		todos[i] = todo
		i++
	}

	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.Store[strconv.Itoa(todo.ID)] = todo
	json.NewEncoder(w).Encode(todo)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todo, ok := models.Store[vars["id"]]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := models.Store[vars["id"]]; !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	delete(models.Store, vars["id"])
	w.WriteHeader(http.StatusNoContent)
}
