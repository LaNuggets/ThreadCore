package main

import (
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search/" {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/search.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	search := r.URL.Query().Get("q")
	media := r.URL.Query().Get("media") // media options  : posts, communities, comments , users
	sort := r.URL.Query().Get("sort")
	time := r.URL.Query().Get("time")

	searchedPost := database.GetPostsBySearchString(search)

	// postContent := api.DisplayPosts(searchedPost)

	searchPage := struct {
		Search       string
		Media        string
		Sort         string
		Time         string
		SearchedPost []database.PostDisplay
	}{
		Search:       search,
		Media:        media,
		Sort:         sort,
		Time:         time,
		SearchedPost: searchedPost,
	}

	err = tmpl.Execute(w, searchPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
