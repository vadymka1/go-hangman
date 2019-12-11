package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/go-hangman/services"
	"log"
	"net/http"
)

func main()  {
	fmt.Println("Starting ...")

	port := ":9090"

	router := mux.NewRouter()

	router.HandleFunc("/", services.ShowForm).Methods("GET")
	router.HandleFunc("/hangman", services.CheckData).Methods("POST")

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}

}

