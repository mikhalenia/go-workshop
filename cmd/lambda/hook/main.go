package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("REQUEST: %v", r)
		fmt.Fprintf(w, "Hello, %q", "PONG")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}