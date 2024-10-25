package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// HomeHandler serves the main page and handles partial updates.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Serve the full page
		tmplPath := filepath.Join("views", "base.html")
		indexPath := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(tmplPath, indexPath)
		if err != nil {
			http.Error(w, "Error loading page", http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			fmt.Println(tmplPath)
			fmt.Println(err)
			http.Error(w, "Error executing template", http.StatusInternalServerError)

		}
	} else if r.Method == http.MethodPost && r.URL.Path == "/partial" {
		// Serve partial content
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, "<p>This content was loaded via HTMX!</p>")
	}
}
