package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// Restful Apis
func startServer() {
	fmt.Println("Starting Server ...... ")

	router := chi.NewRouter()
	router.Get("/api/searchAll",searchAllHandler)
	router.Get("/api/searchBySource",searchBySourceHandler)
	router.Get("/api/searchByTitle",searchByTitleHandler)
	router.Get("/api/searchByBody",searchByBodyHandler)

	log.Fatal(http.ListenAndServe(":8080",router))
	fmt.Println("Server is listening on port 8080...")

}

func searchAllHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(SearchAll())
}

func searchBySourceHandler(w http.ResponseWriter, r *http.Request) {
	source := r.FormValue("source")
	json.NewEncoder(w).Encode(SearchBySource(source))
}

func searchByTitleHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	json.NewEncoder(w).Encode(SearchByTitle(title))
}

func searchByBodyHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	json.NewEncoder(w).Encode(SearchByBody(body))
}
