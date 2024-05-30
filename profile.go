package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func User(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/user/" {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/profile.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	username := strings.ReplaceAll(r.URL.Path, "/user/", "")
	if strings.Contains(username, "/") {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
	}

	userPage := struct {
		Name string
	}{
		Name: username,
	}

	err = tmpl.Execute(w, userPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
