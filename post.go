package main

import (
	"ThreadCore/database"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	postUuid := strings.ReplaceAll(r.URL.Path, "/post/", "")
	if strings.Contains(postUuid, "/") {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
	}
	post := database.GetPostByUuid(postUuid)
	fmt.Println(post)

	postPage := struct {
		Post database.PostInfo
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
