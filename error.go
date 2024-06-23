package main

import (
	"html/template"
	"log"
	"net/http"
)

// !The Error function is used to create the error page which appears when the serveur encouter an error. She take as arguments a writer and a request.
func Error(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./templates/error.html")
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
