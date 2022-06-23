package handler

import (
	"encoding/json"
	"fmt"
	"go-restful/pkg/data"
	"go-restful/pkg/dto"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("xyz")
	result, err := data.GetAllData()
	if err != nil {
		fmt.Println(err)
	}
	// result := make([]dto.Todo, 0)
	// // for _, item := range r{
	// // 	result = append(result, dto.Todo{
	// // 		ID: r.,
	// // 	})
	// // }
	responseWithJson(writer, http.StatusCreated, result)
}

func GetTodoById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}
	result, err := data.GetDataByID(int64(id))
	if err != nil {
		fmt.Println(err)
	}
	responseWithJson(writer, http.StatusCreated, result)

}

func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var newTodo dto.Todo
	fmt.Println("abc")
	if err := json.NewDecoder(request.Body).Decode(&newTodo); err != nil {
		fmt.Println(newTodo)
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	_, err := data.CreateData(newTodo)
	if err != nil {
		fmt.Println(err)
	}
	responseWithJson(writer, http.StatusCreated, map[string]string{"message": "Create Successfully"})
}

func UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	var updateTodo dto.Todo
	if err := json.NewDecoder(request.Body).Decode(&updateTodo); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	_, err = data.UpdateDataByID(id, bson.D{
		{"$set", bson.D{
			{"id", updateTodo.ID},
			{"name", updateTodo.Name},
			{"content", updateTodo.Content},
			{"status", updateTodo.Status},
		}},
	})
	if err != nil {
		fmt.Println(err)
	}
	responseWithJson(writer, http.StatusCreated, map[string]string{"message": "Update Successfully"})
}

func DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}
	_, err = data.DeleteDataByID(id)
	if err != nil {
		fmt.Println(err)
	}
	responseWithJson(writer, http.StatusCreated, map[string]string{"message": "Delete Successfully"})
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}
