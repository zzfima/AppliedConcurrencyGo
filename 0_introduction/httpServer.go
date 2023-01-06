package main

import (
	"log"
	"net/http"
)

func mainHTTPServer() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello")
}
