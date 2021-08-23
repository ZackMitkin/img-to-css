package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	log.Println("Listening on http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
