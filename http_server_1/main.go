package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
		fmt.Printf("%v\n", r.Method)
	})
	http.ListenAndServe(":8771", nil)
}
