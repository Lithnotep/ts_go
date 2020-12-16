package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/person", personHandler)
	http.HandleFunc("/people", peopleHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	response := response{Message: "Max's API"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("oops")
	}

	fmt.Fprint(w, string(data))
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	response := response{Message: Conn()}
	data, err := json.Marshal(response)
	if err != nil {
		panic("oops")
	}

	fmt.Fprint(w, string(data))
}
