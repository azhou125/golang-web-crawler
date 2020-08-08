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
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(SearchAll()), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}

func searchBySourceHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	source := query.Get("source")

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(SearchBySource(source)), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}

func searchByTitleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := query.Get("title")

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(SearchByTitle(title)), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}

func searchByBodyHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	body := query.Get("body")

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(SearchByBody(body)), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}
