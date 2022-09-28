package main

import (
	"fmt"
	function "go-postgres/functions"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type movie struct {
	gorm.Model
	MovieId   string `json:"movieid"`
	MovieName string `json:"moviename"`
	Isactive  bool   `json:"Isactive"`
}

var (
	drivers = []movie{
		{MovieName: "abc", MovieId: "abc", Isactive: true},
	}
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "sanjai"
	dbname   = "postgres"
)

func InitialMigration() {
	//router := mux.NewRouter()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")

	}
	db.AutoMigrate(&movie{})
	for index := range drivers {
		db.Create(&drivers[index])
	}
}

func initializeRouter() int {

	r := mux.NewRouter()
	r.HandleFunc("/", function.Getmovie).Methods("GET")
	r.HandleFunc("/movie", function.Createmovie).Methods("POST")
	r.HandleFunc("/movie/{id}", function.Updatemovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", function.Softdelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe("Localhost:5000", r))

	return 0
}

func main() {
	InitialMigration()
	initializeRouter()
}

// Sample function to print a line
func Sample() string {
	out := "failed to connect database"
	return out
}
