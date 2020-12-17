package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "touchsource"
)

var db *sql.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func main() {
	port := 8080

	http.HandleFunc("/person", personHandler)
	http.HandleFunc("/people", peopleHandler)

	log.Printf("Server starting on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func dbConn(q string, w http.ResponseWriter) {
	people, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer people.Close()

	var list []*peopleList
	for people.Next() {
		p := new(peopleList)
		err = people.Scan(&p.ID, &p.First, &p.Last)
		if err != nil {
			// handle this error
			panic(err)
		}
		list = append(list, p)
	}
	if err := people.Err(); err != nil {
		fmt.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
		fmt.Println(err)
	}
	return

}

type peopleList struct {
	First string `json:"first"`
	Last  string `json:"last"`
	ID    int64  `json:"id"`
}

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	ID    int64  `json:"id"`
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	dbConn("SELECT * FROM People ORDER BY last ASC, first ASC;", w)
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	dbConn("SELECT * FROM People ORDER BY last ASC, first ASC;", w)
}
