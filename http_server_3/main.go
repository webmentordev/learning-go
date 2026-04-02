package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age  := r.URL.Query().Get("age")

	if name == "" {
		name = "stranger"
	}

	fmt.Fprintf(w, "Hello, %s! Age: %s\n", name, age)
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}