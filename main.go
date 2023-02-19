package main

import (
	"fmt"
	"log"
	"net/http"
	"test-go-api/handler"

	"github.com/gorilla/mux"
)

func main() {
	// initialize router
	router := mux.NewRouter()

	// Define route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hallo Selamat Datang")
	}).Methods(http.MethodGet)
	router.HandleFunc("/users", handler.GetDataUser).Methods("GET")
	router.HandleFunc("/users/{id}", handler.GetDataByID).Methods("GET")
	router.HandleFunc("/users/", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")

	// Run server
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
