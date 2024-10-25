package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ihce/Oxide-Website.git/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler) // Serve the home page

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
