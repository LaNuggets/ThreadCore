package main

import (
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
	tmpl, err := template.ParseFiles("./templates/connection.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	username, email, password := getIdentifier(r)

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func getIdentifier(r *http.Request) (*string, string, string) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	username := r.FormValue("username")

	if username == "" {
		return nil, email, password
	} else {
		return &username, email, password
	}
}

func creationProfile() {

	// _, err = db.Exec("INSERT INTO cartes_utilisateurs (nom, prenom, date_naissance, moyenne_generale, photo_profile, classe) VALUES (?, ?, ?, ?, ?, ?)", nom, prenom, date, moyenne, photo, classe)
	// if err != nil {
	// 	http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
	// 	return
	// }
}
