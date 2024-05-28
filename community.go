package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Communities(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/community" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("./templates/community.html") // Read the home page
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

func Community(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/community/" {
		http.Redirect(w, r, "/community", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("./templates/community.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	communityName := strings.ReplaceAll(r.URL.Path, "/community/", "")
	if strings.Contains(communityName, "/") {
		http.Redirect(w, r, "/community", http.StatusSeeOther)
	}

	community := struct {
		Name string
	}{
		Name: communityName,
	}

	err = tmpl.Execute(w, community)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
