package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// Home serves the main page.
func Home(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("internal", "views", "base.html")
	tmpl, err := template.ParseFiles(tmplPath, filepath.Join("internal", "views", "index.html"))
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base", nil)
}

// PartialHandler serves HTMX dynamic content.
func PartialHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>This content was loaded via HTMX!</p>")
}
