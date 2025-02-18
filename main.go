package main

import (
	"countries/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.CountryList)
	mux.HandleFunc("GET /country/{name}", handlers.CountryDetail)
	mux.HandleFunc("GET /search", handlers.SearchCountry)

	log.Println("Server starting on :8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
