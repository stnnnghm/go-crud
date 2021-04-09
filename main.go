package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Global Article array to simulate DB
var Articles []Article

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Goodbye", Desc: "Another Description", Content: "More Content"},
	}

	// Could also:
	// http.HandleFunc("/", homePage), http.HandleFunc("articles", returnAllArticles), etc...
	// log.Fatal(http.ListenAndServe(":8080", nil))
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	// Map API endpoints
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
