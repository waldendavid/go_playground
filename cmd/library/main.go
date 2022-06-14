package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waldendavid/restapi/pkg/library"
)

func main() {
	libService := library.NewService()
	libHandler := library.NewHandler(libService)

	r := mux.NewRouter()
	r.StrictSlash(false)

	r.HandleFunc("/api/books", libHandler.GetBooks()).Methods("GET")
	r.HandleFunc("/api/books/{id}", libHandler.GetBook()).Methods("GET")
	r.HandleFunc("/api/books", libHandler.CreateBook()).Methods("POST")
	r.HandleFunc("/api/books/{id}", libHandler.UpdateBook()).Methods("PUT")
	r.HandleFunc("/api/books/{id}", libHandler.DeleteBook()).Methods("DELETE")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())

}

// import (
// 	"encoding/json"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )
// //https://www.youtube.com/watch?v=SonwZ6MF5BE

// // Init Books vas as a slice Book struct
// var books []Book

// //Get All Books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// // Get Single Book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) //Get params
// 	//Loop through books and find with id
// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// // Create a New Book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(10000000)) //Mock Id = not safe
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }

// // Update Book
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) //Get params
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// // Delete Book
// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) //Get params
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// func main() {
// 	//init router
// 	r := mux.NewRouter()

// 	// Mock Data @todo implement DB
// 	books = append(books, Book{ID: "1", Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
// 	books = append(books, Book{ID: "2", Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

// 	//Route Handlers / Endpoints
// 	r.HandleFunc("/api/books", getBooks).Methods("GET")
// 	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
// 	r.HandleFunc("/api/books", createBook).Methods("POST")
// 	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
// 	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
// 	log.Fatal(http.ListenAndServe(":8000", r))
// }
