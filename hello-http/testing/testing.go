package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(("/"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	mux.HandleFunc(("/bar"), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, bar!")
	})

	return mux
}

func main() {
	http.ListenAndServe(":8080", MakeWebHandler())
}
