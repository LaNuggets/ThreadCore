package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("./templates/post.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	fmt.Println(title, content)

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
