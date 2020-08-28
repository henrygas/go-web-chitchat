package main

import (
	"fmt"
	"go-web-chitchat/data"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		return
	}

	_, err = session(w, r)
	if err != nil {
		generateHTML(w, threads, "layout", "public.navbar", "index")
	} else {
		generateHTML(w, threads, "layout", "private.navbar", "index")
	}
}

func err(w http.ResponseWriter, r *http.Request) {

}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(w, "layout", data)
}
