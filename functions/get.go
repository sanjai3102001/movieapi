package function

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB

type movie struct {
	gorm.Model
	MovieId   string `json:"movieid"`
	MovieName string `json:"moviename"`
	Isactive  bool   `json:"Isactive"`
}

// this is a Getmovies function
func Getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var moviess []movie
	db.Find(&moviess)
	json.NewEncoder(w).Encode(moviess)
}
func Getmovieid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movies []movie
	db.Find(&movies)
	json.NewEncoder(w).Encode(movies)
}
