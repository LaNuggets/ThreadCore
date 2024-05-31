package main

import (
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/post/" {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/post.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	postId := strings.ReplaceAll(r.URL.Path, "/post/", "")
	if strings.Contains(postId, "/") {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
	}
	id, _ := strconv.Atoi(postId)
	post := database.GetPostById(id)

	postPage := struct {
		Post database.Post
	}{
		Post: post,
	}

	err = tmpl.Execute(w, postPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
