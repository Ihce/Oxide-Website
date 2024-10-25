package main

import (
	"log"
	"net/http"

	"github.com/Ihce/Oxide-Website.git/internals/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)                  // Serve the main page
	http.HandleFunc("/partial", handlers.PartialHandler) // Handle HTMX partial updates

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
