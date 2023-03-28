package main

import (
	"log"
	"net/http"

	"challenge3/package/database"
	"challenge3/package/handlers"
)

func main() {
	DB := database.Init()
	h := handlers.New(DB)

	// localhost:8000/books
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAllBooks(w, r)
		case http.MethodPost:
			h.AddBook(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	// localhost:8000/books/?id=1
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetBook(w, r)
		case http.MethodPut:
			h.UpdateBook(w, r)
		case http.MethodDelete:
			h.DeleteBook(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	log.Println("API is running!")
	http.ListenAndServe(":8000", nil)
}
