package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	fmt.Println("TODO - REST Service running on 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index is the main function called when / rest endpoint hits.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to TODO REST APP!")
}

// TodoIndex is the main function called when / rest endpoint hits.
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

// TodoShow is the main function called when / rest endpoint hits.
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
