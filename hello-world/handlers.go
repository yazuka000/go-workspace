package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, world!")
	// fmt.Fprintf(w, "This is the home page")

	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))

	renderTemplate(w, "about.page.tmpl")
}
