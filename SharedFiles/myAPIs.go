package SharedFiles

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// Restful Apis
func StartServer() {
	fmt.Println("Starting Server ...... ")

	router := chi.NewRouter()
	router.Get("/api/searchAll",SearchAllHandler)
	router.Get("/api/searchBySource",SearchBySourceHandler)
	router.Get("/api/searchByTitle",SearchByTitleHandler)
	router.Get("/api/searchByBody",SearchByBodyHandler)

	log.Fatal(http.ListenAndServe(":8080",router))
	fmt.Println("Server is listening on port 8080...")

}

func SearchAllHandler(w http.ResponseWriter, r *http.Request) {
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(SearchAll()), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}

func SearchBySourceHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	source := query.Get("source")

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(SearchBySource(source)), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}

func SearchByTitleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := query.Get("title")

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(SearchByTitle(title)), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}

func SearchByBodyHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	body := query.Get("body")

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(SearchByBody(body)), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}
