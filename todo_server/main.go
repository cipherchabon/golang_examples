package main

import (
	"net/http"

	"github.com/cypherchabon/golang_examples/todo_server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/api/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/api/todos/{id}", handlers.GetTodo).Methods("GET")
	r.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
