package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/amulya-leapfrog/go-todo/services"
	htmlGenerator "github.com/amulya-leapfrog/go-todo/static"
	"github.com/amulya-leapfrog/go-todo/structs"
	"github.com/go-chi/chi"
)

func FetchTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := services.FetchTodos()
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	htmlGenerator.HtmlGenerator(w, todos)
}

func FetchById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID, must be a number", http.StatusBadRequest)
		return
	}

	todo, err := services.FetchById(id)
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo structs.CreateTodo
	payload := r.Body

	decoder := json.NewDecoder(payload)
	err := decoder.Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if newTodo.Task == "" {
		http.Error(w, "Task field is required", http.StatusBadRequest)
		return
	}

	if newTodo.Status == "" {
		http.Error(w, "Status field is required", http.StatusBadRequest)
		return
	}

	_, err = services.CreateTodo(newTodo)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}

	successResponse := map[string]string{"message": "Task successfully created"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to create record", http.StatusInternalServerError)
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID, must be a number", http.StatusBadRequest)
		return
	}

	var updateTodo structs.CreateTodo
	payload := r.Body

	decoder := json.NewDecoder(payload)
	err = decoder.Decode(&updateTodo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if updateTodo.Task == "" {
		http.Error(w, "Task field is required", http.StatusBadRequest)
		return
	}

	if updateTodo.Status == "" {
		http.Error(w, "Status field is required", http.StatusBadRequest)
		return
	}

	_, err = services.UpdateTodo(id, updateTodo)
	if err != nil {
		http.Error(w, "Failed to udpate record", http.StatusInternalServerError)
		return
	}

	successResponse := map[string]string{"message": "Task successfully updated"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to update record", http.StatusInternalServerError)
	}
}

func UpdateAllTodo(w http.ResponseWriter, r *http.Request) {
	rowsAffected, err := services.UpdateAllTodo()
	if err != nil {
		http.Error(w, "Failed to udpate records", http.StatusInternalServerError)
		return
	}

	successResponse := map[string]string{
		"message": fmt.Sprintf("%d rows successfully updated", rowsAffected),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to update records", http.StatusInternalServerError)
	}
}

func DeleteAllTodo(w http.ResponseWriter, r *http.Request) {
	rowsAffected, err := services.DeleteAllTodo()
	if err != nil {
		http.Error(w, "Failed to delete records", http.StatusInternalServerError)
		return
	}

	successResponse := map[string]string{
		"message": fmt.Sprintf("%d rows successfully deleted", rowsAffected),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to delete records", http.StatusInternalServerError)
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID, must be a number", http.StatusBadRequest)
		return
	}

	_, err = services.DeleteTodo(id)
	if err != nil {
		http.Error(w, "Failed to delete record", http.StatusInternalServerError)
		return
	}

	successResponse := map[string]string{"message": "Task successfully deleted"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(successResponse); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to delete record", http.StatusInternalServerError)
	}
}
