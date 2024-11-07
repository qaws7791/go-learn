package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    // get the query values
	name := values.Get("name") // get the name value
	if name == "" {
		name = "Guest"
	}

	id, _ := strconv.Atoi(values.Get("id"))             // get the id value
	fmt.Fprintf(w, "Hello %s! Your id is %d", name, id) // write data to response
}

func main() {
	http.HandleFunc("/bar", barHandler) // register the handler function
	http.ListenAndServe(":8080", nil)   // listen on port 8080
}
