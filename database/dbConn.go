package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
    DB *sql.DB
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    hostName := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
    userName := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")

    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
        userName, password, dbName, hostName)

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("=============== Connected to the database successfully! ===============")
}
