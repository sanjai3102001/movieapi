package main

import (
	"go-postgres/function"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", function.CreateItem).Methods("GET")
	router.HandleFunc("/movie/{id}", function.ReadingItemid).Methods("GET")
	router.HandleFunc("/movie", function.CreateItem).Methods("POST")
	router.HandleFunc("/movie/{id}", function.UpdateItems).Methods("PUT")
	router.HandleFunc("/movie/{id}", function.Softdelete).Methods("DELETE")
	return router
}

func TestCreateItem(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestReadItem(t *testing.T) {
	request, _ := http.NewRequest("GET", "/movie/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
