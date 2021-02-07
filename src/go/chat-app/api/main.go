package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article represents an article for our API
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is the list of all articles
var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit : returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint hit : homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
		{Title: "hello", Desc: "Article Description", Content: "Article content"},
		{Title: "hello2", Desc: "Article 2 Description", Content: "Article 2 content"},
	}
	handleRequests()
}
