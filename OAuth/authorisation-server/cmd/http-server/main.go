package main

import (
	"log"
	"net/http"

	"github.com/7om3k-dev/OAuth/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /authorise", handlers.Authorise)

	log.Print("starting auth server")
	log.Fatal(http.ListenAndServe(":3001", mux))
}
