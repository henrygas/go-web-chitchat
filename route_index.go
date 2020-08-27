package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		return
	}

	_, err = session(w, r)

	publicFiles := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html",
	}
	privateFiles := []string{
		"templates/layout.html",
		"templates/private.navbar.html",
		"templates/index.html",
	}
	var templates *template.Template
	if err != nil {
		templates = template.Must(template.ParseFiles(publicFiles...))
	} else {
		templates = template.Must(template.ParseFiles(privateFiles...))
	}
	_ = templates.ExecuteTemplate(w, "layout", threads)
}

func err(w http.ResponseWriter, r *http.Request) {

}

