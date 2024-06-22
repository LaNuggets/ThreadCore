package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
)

// !The Home function is used to create the home page which is the main page of the forums. She take as arguments a writer and a request.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("./templates/home.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}
	user_uuid := api.GetCookie("uuid", r)
	user := database.GetUserByUuid(user_uuid, w, r)

	userUuid := api.GetCookie("uuid", r)
	userProfile := database.GetUserByUuid(userUuid, w, r).Profile

	homePage := struct {
		Connected bool
		Profile   string
		Username  string
	}{
		Connected: userUuid != "",
		Profile:   userProfile,
		Username:  user.Username,
	}
	err = tmpl.Execute(w, homePage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
