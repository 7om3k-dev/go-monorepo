package main

import (
	"log"
	"net/http"

	"github.com/7om3k-dev/OAuth/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /authorise", handlers.Authorise)

	if err := http.ListenAndServe(":3001", mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
