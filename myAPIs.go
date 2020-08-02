package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/olivere/elastic/v7"

	"github.com/go-chi/chi"
)

// Restful Apis
func startServer() {
	fmt.Println("Starting Server ...... ")
	router := chi.NewRouter()
	router.Get("/api/getExample",getHandler)
	router.Post("/api/postExample",postHandler)

	log.Fatal(http.ListenAndServe(":8080",router))
	fmt.Println("Server is listening on port 8080...")

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You got me")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You just send me a post req!")
}


// 业务逻辑
func getESClient() (*elastic.Client) {

	client, err :=  elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	fmt.Println("ES initialized...")

	return client

}

func insertNews(esclient *elastic.Client, data StdNew){
	ctx := context.Background()
	dataJSON, err := json.Marshal(data)
	js := string(dataJSON)
	_, err = esclient.Index().
		Index("news").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}

//type Student struct {
//	Name         string  `json:"name"`
//	Age          int64   `json:"age"`
//	AverageScore float64 `json:"average_score"`
//}

//func insertData() {
//
//	ctx := context.Background()
//	esclient, err := GetESClient()
//
//
//	//creating student object
//	newStudent := Student{
//		Name:         "Gopher doe",
//		Age:          10,
//		AverageScore: 99.9,
//	}
//
//	dataJSON, err := json.Marshal(newStudent)
//	js := string(dataJSON)
//	_, err = esclient.Index().
//		Index("students").
//		BodyJson(js).
//		Do(ctx)
//
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
//
//}
//
//
//func searchData() {
//
//	ctx := context.Background()
//	esclient, err := GetESClient()
//	if err != nil {
//		fmt.Println("Error initializing : ", err)
//		panic("Client fail ")
//	}
//
//	var students []Student
//
//	searchSource := elastic.NewSearchSource()
//	searchSource.Query(elastic.NewMatchQuery("name", "Doe"))
//
//	/* this block will basically print out the es query */
//	queryStr, err1 := searchSource.Source()
//	queryJs, err2 := json.Marshal(queryStr)
//
//	if err1 != nil || err2 != nil {
//		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
//	}
//	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
//	/* until this block */
//
//	searchService := esclient.Search().Index("students").SearchSource(searchSource)
//
//	searchResult, err := searchService.Do(ctx)
//	if err != nil {
//		fmt.Println("[ProductsES][GetPIds]Error=", err)
//		return
//	}
//
//	for _, hit := range searchResult.Hits.Hits {
//		var student Student
//		err := json.Unmarshal(hit.Source, &student)
//		if err != nil {
//			fmt.Println("[Getting Students][Unmarshal] Err=", err)
//		}
//
//		students = append(students, student)
//	}
//
//	if err != nil {
//		fmt.Println("Fetching student fail: ", err)
//	} else {
//		for _, s := range students {
//			fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
//		}
//	}
//
//}
