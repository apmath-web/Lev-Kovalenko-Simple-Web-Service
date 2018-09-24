package main

import (
	"Lev-Kovalenko-Simple-Web-Service/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Hello_world)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
