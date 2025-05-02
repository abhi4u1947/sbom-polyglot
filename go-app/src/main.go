package main

import (
	"log"
	"net/http"

	"github.com/nc/sbom-polyglot/go-app/src/handlers"
)

func main() {
	router := handlers.NewRouter()
	
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
} 