package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book struct(Model)
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Isbn   string  `json:"isbn"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init book var as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("w", w)
	json.NewEncoder(w).Encode(books)

}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "applicaton/json")
	params := mux.Vars(r) // Get any params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "applicaton/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update the book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "applicaton/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}
	json.NewEncoder(w).Encode(books)
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "applicaton/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init the mux router
	r := mux.NewRouter()

	// Mock Data @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "78777", Title: "Book 1", Author: &Author{ID: "1", FirstName: "Sumit", LastName: "Sapkota"}})
	books = append(books, Book{ID: "2", Isbn: "455566", Title: "Book 2", Author: &Author{ID: "2", FirstName: "Saroj", LastName: "Shakya"}})

	// Route handles / Endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
