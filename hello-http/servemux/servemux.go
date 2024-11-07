package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // Create a new ServeMux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About us")
	})

	http.ListenAndServe(":8080", mux)
}
