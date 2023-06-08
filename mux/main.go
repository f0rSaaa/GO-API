package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Album struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}

func getAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var alb []Album
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Println(err)
	}
	queryRes, err := db.Query("select * from albums")
	if err != nil {
		// log.Println("All Post 1")
		// fmt.Println(err)
		log.Println(err)
	}

	defer queryRes.Close()

	for queryRes.Next() {
		var a Album
		err := queryRes.Scan(&a.UserId, &a.Id, &a.Title)
		if err != nil {
			// log.Println("All Posts 2")
			// fmt.Println(err)
			log.Println(err)
		}
		alb = append(alb, a)
	}

	json.NewEncoder(w).Encode(alb)
}

func getUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Println(err)
	}
	result, err := db.Query("SELECT * FROM albums WHERE user_id = ?", params["id"])
	if err != nil {
		// log.Println("Single Post 1")
		// fmt.Println(err)
		log.Println(err)
	}
	defer result.Close()
	var alb []Album
	for result.Next() {
		var a Album
		err := result.Scan(&a.UserId, &a.Id, &a.Title)
		if err != nil {
			// log.Println("Single Post 2")
			// fmt.Println(err)
			log.Println(err)
		}
		alb = append(alb, a)
	}
	json.NewEncoder(w).Encode(alb)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/getAllData", getAllData).Methods("GET")
	router.HandleFunc("/getUserData/{id}", getUserData).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router), nil)

}
