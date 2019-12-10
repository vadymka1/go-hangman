package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/go-hangman/services"
	"log"
	"net/http"
)

var word, port string

var error, rightWord int


func main()  {
	fmt.Println("Starting ...")

	router := mux.NewRouter()

	router.HandleFunc("/", services.ShowForm).Methods("GET")
	router.HandleFunc("/hangman", services.CheckData).Methods("POST")

	if err := http.ListenAndServe(":9090", router); err != nil {
		log.Fatal(err)
	}

}

