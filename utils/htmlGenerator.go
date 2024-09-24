package htmlGenerator

import (
	"html/template"
	"log"
	"net/http"

	"github.com/amulya-leapfrog/go-todo/structs"
)

func HtmlGenerator(w http.ResponseWriter, todos []structs.Todo) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err = t.Execute(w, todos)
	if err != nil {
		http.Error(w, "Unable to generate HTML", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}
