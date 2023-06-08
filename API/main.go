package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type Albums struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}

//title url: https://jsonplaceholder.typicode.com/albums
//albums url: https://jsonplaceholder.typicode.com/photos?albumId=1

func main() {
	fmt.Println("GO Api")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/albums")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(content))

	// data := string(content)
	// fmt.Println(reflect.TypeOf(data))

	// jsonData := []byte()
	// fmt.Println(jsonData)

	var alb []Albums
	if err := json.Unmarshal(content, &alb); err != nil { // Parse []byte to go struct pointer
		// fmt.Println(err)
		fmt.Println("Can not unmarshal JSON")
	}

	// fmt.Println(result)
	fmt.Println(reflect.TypeOf(alb))
	albumLenght := len(alb)
	fmt.Println("length of data", albumLenght)

	for _, album := range alb {
		fmt.Println("The UserId is: ", album.UserId)
		fmt.Println("The AlbumId is: ", album.Id)
		fmt.Println("The AlbumName is: ", album.Title)
		fmt.Println()
	}

	fmt.Println("Connecting to database")

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(db.Stats())
	defer db.Close()

	//insserting into database

	for _, album := range alb {
		// fmt.Println("The UserId is: ", album.UserId, reflect.TypeOf(album.UserId))
		// fmt.Println("The AlbumId is: ", album.Id, reflect.TypeOf(album.Id))
		// fmt.Println("The AlbumName is: ", album.Title, reflect.TypeOf(album.Title))
		// fmt.Println()
		// break
		_, err := db.Query("insert into albums(user_id, id, title) values(?,?,?)", album.UserId, album.Id, album.Title)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Data inserted into table", len(alb))
}
