package handlers

import (
	"fmt"
	"net/http"
)

func PartialHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>This content was loaded via HTMX!</p>")
}
