package routes

import (
	"net/http"

	"github.com/amulya-leapfrog/go-todo/controllers"

	"github.com/go-chi/chi"
)

func TodoHandler() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", controllers.FetchTodos)
		r.Get("/{id}", controllers.FetchById)
		r.Post("/", controllers.CreateTodo)
		r.Put("/{id}", controllers.UpdateTodo)
		r.Patch("/", controllers.UpdateAllTodo)
		r.Delete("/{id}", controllers.DeleteTodo)
		r.Delete("/", controllers.DeleteAllTodo)
	})
	return rg
}
