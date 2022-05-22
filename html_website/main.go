package main

import (
	// package for printing
	"fmt"
	// package for http stuff
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get away from my home!")
}
