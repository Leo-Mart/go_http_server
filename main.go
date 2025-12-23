package main

import (
	"log"
	"net/http"
)

func main() {
	const filePathRoot = "."
	const port = "8080"
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(filePathRoot)))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving file on port: %s", port)
	log.Fatal(server.ListenAndServe())
}
