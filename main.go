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

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)

	fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}
}
