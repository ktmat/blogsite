package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the posts page!")
})
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
