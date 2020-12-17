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

func main() {
	port := 8080

	http.HandleFunc("/person", personHandler)
	http.HandleFunc("/people", peopleHandler)

	log.Printf("Server starting on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func dbConn(q string, w http.ResponseWriter) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	people, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer people.Close()

	var arr [2]string
	var list [][2]string
	for people.Next() {
		err = people.Scan(&arr[0], &arr[1])
		if err != nil {
			// handle this error
			panic(err)
		}
		list = append(list, arr)
	}
	if err := people.Err(); err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(list)
	fmt.Fprint(w, string(data)
	return

}

// type peopleList struct {
// 	Person
// }

func personHandler(w http.ResponseWriter, r *http.Request) {
	dbConn("SELECT first, last FROM People ORDER BY last ASC, first ASC;", w)
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	dbConn("SELECT first, last FROM People ORDER BY last ASC, first ASC;", w)
}
