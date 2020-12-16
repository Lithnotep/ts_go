package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=millmer dbname=touchsource sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// first := "Donnell"
	// last := "Abrahamson"
	rows, err := db.Query("SELECT * FROM People;")
	fmt.Println(rows)
}
