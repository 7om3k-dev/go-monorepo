package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /authorise", authorize)
	mux.HandleFunc("POST /approve", approve)

	srv := http.Server{
		Addr:    ":8099",
		Handler: mux,
	}

	log.Print("starting auth server")
	log.Fatal(srv.ListenAndServe())
}
