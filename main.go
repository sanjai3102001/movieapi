package main

import (
	"go-postgres/function"
	"log"

	//"main/meathods"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// CreateTablee()
	r.HandleFunc("/", function.CreateItem).Methods("GET")
	r.HandleFunc("/movie/{id}", function.ReadingItemid).Methods("GET")
	r.HandleFunc("/movie", function.CreateItem).Methods("POST")
	r.HandleFunc("/movie/{id}", function.UpdateItems).Methods("PUT")
	r.HandleFunc("/movie/{id}", function.Softdelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe("Localhost:5000", r))

}
