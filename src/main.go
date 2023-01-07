package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("GET / ")
	})

	http.HandleFunc("/todos", func(res http.ResponseWriter, req *http.Request) {
		log.Println("GET /todos ")
		t, err := template.ParseFiles("src/todos.html")

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			log.Println("Template Parsing Error", err)
		}

		t.Execute(res, nil)

	})

	PORT := 5005
	log.Printf("Server is running at port :%d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))

}
