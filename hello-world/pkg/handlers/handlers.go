package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, world!")
	// fmt.Fprintf(w, "This is the home page")

	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))

	render.RenderTemplate(w, "about.page.tmpl")
}
