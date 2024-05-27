package main

import (
	"html/template"
	"log"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./templates/notFound.html")
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
