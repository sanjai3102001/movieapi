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
	r.HandleFunc("/movie/1", function.ReadingItemid).Methods("GET")
	r.HandleFunc("/movie", function.CreateItem).Methods("POST")
	r.HandleFunc("/movie/2", function.UpdateItems).Methods("PUT")
	r.HandleFunc("/movie/2", function.Softdelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe("Localhost:5000", r))

}
