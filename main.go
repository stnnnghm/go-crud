package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Global Article array to simulate DB
var Articles []Article

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Router")

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Goodbye", Desc: "Another Description", Content: "More Content"},
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

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	// Grab id from url (/article/{id})
	vars := mux.Vars(r)
	id := vars["id"]

	// Loop over articles
	// if article.Id equals the key passed in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func handleRequests() {
	// gorilla mux router
	// create new instance of mux router
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/articles", returnAllArticles)
	r.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", r))

	// net/http router
	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", returnAllArticles)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
