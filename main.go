package main

import (
	"fmt"
	"log"
	"net/http"

	"test-go-api/handler"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello and welcome.")
	})

	http.HandleFunc("/user", handler.UserHandler)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
