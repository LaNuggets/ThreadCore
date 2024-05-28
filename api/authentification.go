package api

import (
	"ThreadCore/database"
	"net/http"
)

func Authentication(w http.ResponseWriter, r *http.Request) string {
	username, email, password := GetIdentifier(r)
	if email != "" {
		ChooseConnectionOrCreation(username, email, password, w, r)
	}
	return *&email
}

func GetIdentifier(r *http.Request) (*string, string, string) {
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
		ConnectionProfile(email, password, w, r)
	} else {
		CreationProfile(*username, email, password)
	}
}

func CreationProfile(username string, email string, password string) {
	user := AddUserValue(username, email, password)
	database.AddUser(user)
	println(user.Username)
	println(user.Email)
	println(user.Password)
	println("Creation successful")
}

func ConnectionProfile(email string, password string, w http.ResponseWriter, r *http.Request) {
	user := database.GetUserByEmail(email)
	if user.Password == password {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		println("Welcome", user.Username)
		//Todo: Cookie
	} else {
		http.Redirect(w, r, "/connection?error=password_taken", http.StatusFound)
	}

}

func AddUserValue(username string, email string, password string) database.User {
	user := database.User{nil, "azrar-7894-d5f5d", "picture", email, username, password}
	return user
}
