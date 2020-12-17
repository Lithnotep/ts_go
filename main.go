package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

	http.HandleFunc("/person/", personHandler)
	http.HandleFunc("/people", peopleHandler)

	log.Printf("Server starting on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type personData struct {
	ID    int64  `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/person/")
	fullname := strings.SplitN(name, "/", 2)
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT * FROM People WHERE first = $1 AND last = $2;`

	person := db.QueryRow(sqlStatement, fullname[1], fullname[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	p := new(personData)
	person.Scan(&p.ID, &p.First, &p.Last)

	data, err := json.Marshal(p)
	fmt.Fprint(w, string(data))

}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	people, err := db.Query("SELECT first, last FROM People ORDER BY last ASC, first ASC;")
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
			panic(err)
		}
		list = append(list, arr)
	}
	if err := people.Err(); err != nil {
		fmt.Println(err)
		return
	}

	data, err := json.Marshal(list)
	fmt.Fprint(w, string(data))
}
