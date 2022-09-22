package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, world!")
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// プライベートな関数は、小文字から定義する？
// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println((fmt.Sprintf("Starting application on port %s", portNumber)))
	_ = http.ListenAndServe(portNumber, nil)
}
