package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Community(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/community/" {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
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
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
	}

	userUuid := api.GetCookie("uuid", r)

	communityPage := struct {
		Community database.Community
		Connected bool
	}{
		Community: database.GetCommunityByName(communityName),
		Connected: userUuid != "",
	}

	err = tmpl.Execute(w, communityPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
