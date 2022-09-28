package function

import (
	"encoding/json"
	"net/http"
)

func Createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movies movie
	json.NewDecoder(r.Body).Decode(&movies)
	// fmt.Println(json.NewDecoder(r.Body).Decode(&user))
	db.Create(&movies)
	json.NewEncoder(w).Encode(movies)

}
