package main

import (
	"Lev-Kovalenko-Simple-Web-Service/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HelloWorld)
	http.HandleFunc("/date", handlers.Date)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
