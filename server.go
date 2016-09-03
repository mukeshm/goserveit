package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the server on port 8080")
	router := newRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
