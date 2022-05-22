package main

import (
	// package for printing
	"fmt"
	// package for http stuff
	"net/http"
	// used for routes
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")

	http.ListenAndServe(":5000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Fprintf(w, "Get away from my home!")
}
