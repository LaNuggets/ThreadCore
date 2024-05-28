package main

import (
	"ThreadCore/api"
	"html/template"
	"log"
	"net/http"
)

func Connection(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/connection" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	errorMessage := ""
	if r.URL.Query().Get("error") == "password_taken" {
		errorMessage = "Wrong password !"
	}

	tmpl, err := template.ParseFiles("./templates/connection.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	email := api.Authentication(w, r)

	data := struct {
		ErrorMessage string
		Email        string
	}{
		ErrorMessage: errorMessage,
		Email:        email,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
