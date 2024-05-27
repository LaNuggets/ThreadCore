package main

import (
	"ThreadCore/database"
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
	if r.URL.Query().Get("error") == "username_taken" {
		errorMessage = "This username is already used. Choose an other please."
	}

	tmpl, err := template.ParseFiles("./templates/connection.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	username, email, password := getIdentifier(r)
	ChooseConnectionOrCreation(username, email, password)

	data := struct {
		ErrorMessage string
	}{
		ErrorMessage: errorMessage,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func getIdentifier(r *http.Request) (*string, string, string) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username == "" {
		return nil, email, password
	} else {
		return &username, email, password
	}
}

func ChooseConnectionOrCreation(username *string, email string, password string) {
	if username == nil {
		connectionProfile(email, password)
	} else {
		creationProfile(*username, email, password)
	}
}

func creationProfile(email string, password string, username string) {
	user := addUserValue(username, email, password)
	database.AddUser(database.DB, user)

}

func connectionProfile(email string, password string) {

}

func addUserValue(username string, email string, password string) database.User {
	user := database.User{_, "picture", email, username, password}
	return user
}
