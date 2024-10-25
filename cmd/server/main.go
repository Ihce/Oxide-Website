package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Ihce/Oxide-Website/internals/handlers"
)

// Serve home page
func home(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("internal", "views", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)                           // Root route for main page
	http.HandleFunc("/partial", handlers.PartialHandler) // Route for HTMX partial content

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
