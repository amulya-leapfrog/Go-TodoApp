package main

import (
	"log"
	"net/http"
	"os"

	"github.com/amulya-leapfrog/go-todo/controllers"
	"github.com/amulya-leapfrog/go-todo/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

var (
	serverPort string
)

func main() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	router.Get("/", controllers.FetchTodos)
	router.Mount("/todo", routes.TodoHandler())

	serverPort = os.Getenv("PORT")

	server := &http.Server{
		Handler: router,
		Addr:    ":" + serverPort,
	}

	log.Printf("=============== Server starting on port: %v ===============", serverPort)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
