package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/search.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	search := strings.ReplaceAll(r.URL.Path, "/search/", "")
	if strings.Contains(search, "/") {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
	}

	if r.Method == "POST" {
		media := r.FormValue("type") // media options  : posts, communities, comments ,
		sort := r.FormValue("sort")  // sort options : popular (most likes), recent,
		time := r.FormValue("time")  // time options : all time, year, month, week, day, hour

		fmt.Println(search, media, sort, time)
	}

	searchPage := struct {
		Name string
	}{
		Name: search,
	}

	err = tmpl.Execute(w, searchPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
