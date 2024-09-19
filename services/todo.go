package services

import (
	"log"

	"github.com/amulya-leapfrog/go-todo/models"
	"github.com/amulya-leapfrog/go-todo/structs"
)

func FetchTodos() ([]structs.Todo, error) {
	rows, err := models.FetchTodos()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoList []structs.Todo
	for rows.Next() {
		var todo structs.Todo
		if err := rows.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		todoList = append(todoList, todo)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return todoList, nil
}

func FetchById(id int) (structs.Todo, error) {
	row, err := models.FetchById(id)
	if err != nil {
		log.Println("Record not found: ", err)
		return structs.Todo{}, err
	}
	return row, nil
}

func CreateTodo(data structs.CreateTodo) (int, error) {
	id, err := models.CreateTodo(data)
	if err != nil {
		log.Println("Error creating todo:", err)
		return 0, err
	}
	return id, nil
}

func UpdateTodo(id int, data structs.CreateTodo) (int, error) {
	_, err := models.FetchById(id)
	if err != nil {
		log.Printf("Record with the Id: %v not found", id)
		return 0, err
	}

	_, err = models.UpdateTodo(id, data)
	if err != nil {
		log.Println("Error Updating the given todo: ", err)
		return 0, err
	}

	return 1, nil
}

func UpdateAllTodo() (int, error) {
	rowsAffected, err := models.UpdateAllTodo()
	if err != nil {
		log.Println("Error Updating the todos: ", err)
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteAllTodo() (int, error) {
	rowsAffected, err := models.DeleteAllTodo()
	if err != nil {
		log.Println("Error Deleting the todos: ", err)
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteTodo(id int) (int, error) {
	_, err := models.FetchById(id)
	if err != nil {
		log.Printf("Record with the Id: %v not found", id)
		return 0, err
	}

	err = models.DeleteTodo(id)
	if err != nil {
		log.Println("Error Deleting the given todo: ", err)
		return 0, err
	}

	return 1, nil
}
