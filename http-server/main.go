package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "MOHANAPRABHU D")
	})
	http.HandleFunc("/rollno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212IT181")
	})
	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "IT")
	})
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "purple")
	})

	fmt.Printf("Server running (port=3000), route: http://localhost:3000/helloworld\n")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
