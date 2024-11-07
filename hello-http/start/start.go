package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!") // write data to response
	})

	http.ListenAndServe(":8080", nil) // listen on port 8080
}
