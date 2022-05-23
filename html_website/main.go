package main

import (
	// package for printing
	"fmt"
	// package for http stuff
	"net/http"
	// used for routes
	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	
	// declare the directory filled with static files
	staticFileDirectory := http.Dir("./assets/")
	
	// create a handler that does a /assets/randomAsset.html GET request
	// need to strip prefix so it doesnt search for /assets/assets/file.html
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// sticks the static file handler under the /assets/ page
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":5000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
