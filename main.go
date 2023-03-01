package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

		fmt.Printf("Number of bytes written: %d\n", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
