package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var books []Book

func main() {
	mockBookData()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/book", getBooks).Methods("GET")
	r.HandleFunc("/api/v1/book/{id}", getBookById).Methods("GET")
	r.HandleFunc("/api/v1/book", insertBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func mockBookData() {
	books = append(books, Book{Id: "1", Name: "Deep work"})
	books = append(books, Book{Id: "2", Name: "The pragmatic programmer"})
	books = append(books, Book{Id: "3", Name: "Clean code"})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func insertBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
