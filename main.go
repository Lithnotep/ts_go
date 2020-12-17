package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "touchsource"
)

// func main() {
// 	port := 8080

// 	http.HandleFunc("/person", personHandler)
// 	http.HandleFunc("/people", peopleHandler)

// 	log.Printf("Server starting on port %v\n", 8080)
// 	log.Printf(dbConn, 8080)

// 	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
// }

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	people, err := db.Query("SELECT * FROM People;")

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.Close()
	for people.Next() {
		var first string
		var last string
		var id int64
		err = people.Scan(&id, &first, &last)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(last, first)
	}
}

type response struct {
	Message string `json:"message"`
}

// func personHandler(w http.ResponseWriter, r *http.Request) {
// 	response := response{Message: "Max's API"}
// 	data, err := json.Marshal(response)
// 	if err != nil {
// 		panic("oops")
// 	}

// 	fmt.Fprint(w, string(data))
// }

// func peopleHandler(w http.ResponseWriter, r *http.Request) {

// 	// response := response{Message: "high"}
// 	// data, err := json.Marshal(response)
// 	// if err != nil {
// 	// 	panic("oops")
// 	// }

// 	fmt.Fprint(w, dbConn)

// }
