// package testings
package main

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// var db *gorm.DB
// var err error

// type movie struct {
// 	gorm.Model
// 	MovieId   string `json:"movieid"`
// 	MovieName string `json:"moviename"`
// 	Isactive  bool   `json:"Isactive"`
// }

// var (
// 	drivers = []movie{
// 		{MovieName: "abc", MovieId: "abc", Isactive: true},
// 	}
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "sanjai"
// 	dbname   = "postgres"
// )

// func InitialMigration() int {
// 	//router := mux.NewRouter()
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err = gorm.Open("postgres", psqlInfo)
// 	if err != nil {
// 		// panic("failed to connect database")
// 		panic("0")

// 	}
// 	db.AutoMigrate(&movie{})
// 	for index := range drivers {
// 		db.Create(&drivers[index])
// 	}
// 	return 0
// }

// func initializeRouter() string {

// 	r := mux.NewRouter()
// 	// r.HandleFunc("/", function.Pullm).Methods("GET")
// 	r.HandleFunc("/movie", function.Createmovie).Methods("POST")
// 	r.HandleFunc("/movie/{id}", function.Updatemovie).Methods("PUT")
// 	r.HandleFunc("/movie/{id}", function.Softdelete).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe("Localhost:5000", r))

// 	return "panic: failed to connect database"
// }

// func TestInitializeRouter(t *testing.T) {
// 	expectedoutput := 0
// 	// output := initializeRouter()
// 	output := InitialMigration()
// 	initializeRouter()
// 	if expectedoutput != output {
// 		t.Errorf("Failed ! got %v want %c", output, expectedoutput)
// 	} else {
// 		t.Logf("Success !")
// 	}
// }

func Testunit(t *testing.T) {
	output := initializeRouter()
	expectedoutput := nil

}
