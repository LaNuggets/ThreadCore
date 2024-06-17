package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
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

	cookieUuid := api.GetCookie("uuid", r)

	user := database.GetUserByUuid(cookieUuid, w, r)
	posts := database.GetPostsByUser(user.Id, w, r)

	userPage := struct {
		Connected  bool
		UserBanner string
		UserPp     string
		Username   string
		Posts      []database.PostInfo
	}{
		Connected:  cookieUuid != "",
		UserBanner: user.Banner,
		UserPp:     user.Profile,
		Username:   user.Username,
		Posts:      posts,
	}

	err = tmpl.Execute(w, userPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
