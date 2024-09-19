package models

import (
	"database/sql"
	"log"

	db "github.com/amulya-leapfrog/go-todo/database"
	"github.com/amulya-leapfrog/go-todo/structs"
)

func FetchTodos() (*sql.Rows, error) {
	rows, err := db.DB.Query("SELECT * FROM todo ORDER BY id")
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}
	return rows, nil
}

func FetchById(id int) (structs.Todo, error) {
	findQuery := "SELECT * FROM todo WHERE id = $1"
	var data structs.Todo
	err := db.DB.QueryRow(findQuery, id).Scan(&data.ID, &data.Task, &data.Status, &data.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Record not Found:", id)
		} else {
			log.Println("Error querying database:", err)
		}
		return structs.Todo{}, err
	}
	return data, nil
}

func CreateTodo(data structs.CreateTodo) (int, error) {
	query := "INSERT INTO todo (task, status) VALUES ($1, $2) RETURNING id"

	var id int
	err := db.DB.QueryRow(query, data.Task, data.Status).Scan(&id)
	if err != nil {
		log.Println("Error inserting new todo:", err)
		return 0, err
	}

	return id, nil
}

func UpdateTodo(id int, data structs.CreateTodo) (int, error) {
	query := "UPDATE todo SET task = $1, status = $2 WHERE id = $3"
	_, err := db.DB.Exec(query, data.Task, data.Status, id)
	if err != nil {
		log.Println("Error updating new todo:", err)
		return 0, err
	}
	return 1, nil
}

func UpdateAllTodo() (int, error) {
	query := "UPDATE todo SET status = $1"
	result, err := db.DB.Exec(query, "COMPLETED")
	if err != nil {
		log.Println("Error updating todos:", err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return 0, err
	}

	return int(rowsAffected), nil
}

func DeleteTodo(id int) error {
	query := "DELETE FROM todo WHERE id = $1"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting new todo:", err)
		return err
	}
	return nil
}

func DeleteAllTodo() (int, error) {
	query := "DELETE FROM todo"
	result, err := db.DB.Exec(query)
	if err != nil {
		log.Println("Error deleting todos:", err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return 0, err
	}

	return int(rowsAffected), nil
}
