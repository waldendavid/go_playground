package library

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler(service Library) Handler {
	return &handler{service: service}
}

/*
TODO Handler ma zaiplementowane metody, ale nie ma storingu danych
Ciągle wywoływanie z GetBooks tworzy nowego slice i operacje na nowym??

*/
type Handler interface {
	GetBooks() http.HandlerFunc
	GetBook() http.HandlerFunc
	CreateBook() http.HandlerFunc
	UpdateBook() http.HandlerFunc
	DeleteBook() http.HandlerFunc
}

type handler struct {
	service Library
}

func (h *handler) GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := h.service.GetBooks(r.Context())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
		if err != nil {
			log.Println("Action failed: ", err)
			return
		}

	}
}

// skąd się bierze id? -> z Url : przechowywany w params["id"]
func (h *handler) GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Dupaaa69")
		params := mux.Vars(r)
		id := params["id"]
		b, err := h.service.GetBook(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// respEncoder := json.NewEncoder(w)
		// respEncoder.Encode(b)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(b)

	}
}

func (h *handler) CreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b Book
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		book, err := h.service.CreateBook(r.Context(), b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}

}

func (h *handler) UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		var b Book
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		book, err := h.service.UpdateBook(r.Context(), b, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

func (h *handler) DeleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		err := h.service.DeleteBook(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
