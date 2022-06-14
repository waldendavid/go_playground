package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/provisions/{id}", Provisions)
	http.ListenAndServe(":8080", r)
}

func Provisions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`id := `, id)
	//call http://localhost:8080/provisions/someId in your browser dupa
	//Output : id := someId
}
