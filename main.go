package main

import (
	"fmt"
	"log"
	"net/http"
	"test-go-api/handler"

	"github.com/gorilla/mux"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hallo Selamat Datang")
	//})

	route := mux.NewRouter()
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hallo Selamat Datang")
	}).Methods(http.MethodGet)
	route.HandleFunc("/user/", handler.UserHandler).Methods(http.MethodGet)

	//http.HandleFunc("/buku", handler.BukuHandler)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", route))
}
