package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waldendavid/restapi/pkg/library"
)

func main() {
	// dsn := "host=localhost user=postgres password=secret dbname=postgres port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	libRepository := library.NewRepositoryGorm()
	libService := library.NewServiceLibrary(libRepository)
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
	fmt.Println("Running")
	log.Fatal(srv.ListenAndServe())

}
