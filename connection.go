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
	email := r.FormValue("email")
	password := r.FormValue("password")
	username := r.FormValue("username")

	if username == "" {
		return nil, email, password
	} else {
		return &username, email, password
	}
}

func ChooseCreationOrConnection(email string, password string, username *string) {
	if username == nil {
		connectionProfile(email, password)
	} else {
		creationProfile(email, password, *username)
	}
}

func creationProfile(email string, password string, username string) {
	//  user := addUserStructValue([]user user ,username string, email string, password string)
	// addUser(db, user)
}

func connectionProfile(email string, password string) {

}

// func addUserStructValue([]user user ,username string, email string, password string) []user{
// 	user.username = username
// 	user.email = email
// 	user.password = password
// 	return user
// }
