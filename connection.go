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
	if r.URL.Query().Get("error") == "password_taken" {
		errorMessage = "Wrong password !"
	}

	tmpl, err := template.ParseFiles("./templates/connection.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	username, email, password := getIdentifier(r)
	if email != "" {
		ChooseConnectionOrCreation(username, email, password, w, r)
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
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username == "" {
		return nil, email, password
	} else {
		return &username, email, password
	}
}

func ChooseConnectionOrCreation(username *string, email string, password string, w http.ResponseWriter, r *http.Request) {
	if username == nil {
		println("ici")
		connectionProfile(email, password, w, r)
	} else {
		creationProfile(*username, email, password)
	}
}

func creationProfile(username string, email string, password string) {
	user := addUserValue(username, email, password)
	database.AddUser(user)
	println(user.Username)
	println(user.Email)
	println(user.Password)
	println("Creation successful")
}

func connectionProfile(email string, password string, w http.ResponseWriter, r *http.Request) {
	println("la")
	user := database.GetUserByEmail(email)
	println("laici")
	if user.Password == password {
		println("pas")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		println("Welcome", user.Username)
		//Todo: Cookie
	} else {
		println("pasPas")
		http.Redirect(w, r, "/connection?error=password_taken", http.StatusFound)
	}

}

func addUserValue(username string, email string, password string) database.User {
	user := database.User{"", "picture", email, username, password}
	return user
}
