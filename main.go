package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Book struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var books []Book

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/books", getBooks).Methods("GET")

	mockBookData()

	log.Fatal(http.ListenAndServe(":8000", r))
}

func mockBookData() {
	books = append(books, Book{Id: 1, Name: "Deep work"})
	books = append(books, Book{Id: 2, Name: "The pragmatic programmer"})
	books = append(books, Book{Id: 3, Name: "Clean code"})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
