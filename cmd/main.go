package main

import (
	"fmt"
	"go-restful/pkg/data"
	"go-restful/pkg/handler"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var BaseURI = "mongodb://localhost:27017"

func main() {
	BaseURI = os.Getenv("MONGODB_URI")
	if BaseURI == "" {
		BaseURI = "mongodb://localhost:27017"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	data.ConnectDb(BaseURI)
	fmt.Println("database oke")
	defer data.CloseDb()

	r := mux.NewRouter()

	r.HandleFunc("/api/todo", handler.GetAllTodo).Methods(http.MethodGet)
	r.HandleFunc("/api/todo/{id}", handler.GetTodoById).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handler.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/api/todo/{id}", handler.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/api/todo/{id}", handler.DeleteTodo).Methods(http.MethodDelete)

	http.ListenAndServe(port, r)

}
