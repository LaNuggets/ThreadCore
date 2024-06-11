package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
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

	userUuid := api.GetCookie("uuid", r)
	userProfile := database.GetUserByUuid(userUuid).Profile

	url := strings.ReplaceAll(r.URL.Path, "/post/", "")
	urlQuery := strings.Split(url, "?")
	postUuid := urlQuery[0]
	if strings.Contains(postUuid, "/") {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
	}

	post := database.GetPostByUuid(postUuid)
	comments := database.GetCommentsByPost(post.Id)
	community := database.GetCommunityById(post.Community_id)

	//Time formating
	difference := time.Now().Sub(post.Created)
	postedTime := api.GetFormatedDuration(difference)

	postPage := struct {
		Connected bool
		Profile   string
		Post      database.PostInfo
		PostTime  string
		Community database.Community
		Comments  []database.CommentInfo
	}{
		Connected: userUuid != "",
		Profile:   userProfile,
		Post:      post,
		PostTime:  postedTime,
		Community: community,
		Comments:  comments,
	}

	err = tmpl.Execute(w, postPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
