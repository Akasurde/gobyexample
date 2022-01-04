package main

import (
    "net/http"
    "log"
	"fmt"
	"encoding/json"
    "github.com/gorilla/mux"
)

// Article ...
type Article struct {
    ID     string `json:"Id"`
    Title  string `json:"Title"`
    Author string `json:"author"`
    Link   string `json:"link"`
}
 
type ErrorMsg struct {
    Message string
    ID string
}

// Articles ...
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article/{id}",returnSingleArticle)
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}
 
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    found := false
    for _, article := range Articles {
        if article.ID == key {
            fmt.Println("Article found", article.ID, article.Title)
            json.NewEncoder(w).Encode(article)
            found = true
            break
        }
    }

    if found == false {
        fmt.Println("Article not found")
        json.NewEncoder(w).Encode(ErrorMsg{"No article found", key})
    }
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}
 
func main() {
    Articles = []Article{
        Article{ID: "1",
            Title: "Python Intermediate and Advanced 101",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089KVK23P"},
        Article{ID: "2",
            Title: "R programming Advanced",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089WH12CR"},
        Article{ID: "3",
            Title: "R programming Fundamentals",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089S58WWG"},
    }
    handleRequests()
}