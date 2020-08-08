package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"io/ioutil"
	"net/http"
	"context"
)

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
	fmt.Println(js)
	_, err = esclient.Index().
		Index("news").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}

func SearchAll() string {

	host := "localhost:9200"
	indexPattern := "news"
	url := fmt.Sprintf("http://%s/%s/_search", host, indexPattern)
	queryStr := "{}"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(queryStr)))
	checkError("Failed to fetch news", err)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	checkError("Failed to read news", err)
	return string(body)
}

func SearchBySource(source string) string {

	host := "localhost:9200"
	indexPattern := "news"
	url := fmt.Sprintf("http://%s/%s/_search", host, indexPattern)
	queryStr := fmt.Sprintf(
		`"query":{
			"match":{
				"source":"%s"
			}
		}`,
		source)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(queryStr)))
	checkError("Failed to fetch news", err)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	checkError("Failed to read news", err)
	return string(body)
}

func SearchByTitle(title string) string {

	host := "localhost:9200"
	indexPattern := "news"
	url := fmt.Sprintf("http://%s/%s/_search", host, indexPattern)
	queryStr := fmt.Sprintf(
		`"query":{
			"match":{
				"title":"%s"
			}
		}`,
		title)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(queryStr)))
	checkError("Failed to fetch news", err)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	checkError("Failed to read news", err)
	return string(body)
}

func SearchByBody(newsBody string) string {

	host := "localhost:9200"
	indexPattern := "news"
	url := fmt.Sprintf("http://%s/%s/_search", host, indexPattern)
	queryStr := fmt.Sprintf(
		`"query":{
			"match":{
				"title":"%s"
			}
		}`,
		newsBody)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(queryStr)))
	checkError("Failed to fetch news", err)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	checkError("Failed to read news", err)
	return string(body)
}
