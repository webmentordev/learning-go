package main

import (
	"fmt"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello world #01")
	}

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello world #02")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/hello", h2)

	http.ListenAndServe(":8008", nil)
}
