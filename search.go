package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"fmt"
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
	searchedCommunities := database.GetCommunityBySearchString(search)
	searchUser := database.GetUserByUsername(search)

	var data []string
	switch r.FormValue("media") {
	case "":
		api.NewestPost(searchedPost)
		for i := 0; i < len(searchedPost); i++ {
			fmt.Print("Post: ")
			fmt.Println(searchedPost[i].Content)
			data = append(data, searchedPost[i].Content)
		}
	case "posts":
		api.NewestPost(searchedPost)
		for i := 0; i < len(searchedPost); i++ {
			fmt.Print("Post: ")
			fmt.Println(searchedPost[i].Content)
			data = append(data, searchedPost[i].Content)
		}
	case "communities":
		for i := 0; i < len(searchedCommunities); i++ {
			fmt.Print("Communities:")
			fmt.Println(searchedCommunities[i].Name)
			data = append(data, searchedCommunities[i].Name)
		}

	case "users":
		for i := 0; i < len(searchUser); i++ {
			fmt.Print("User:")
			fmt.Println(searchUser[i].Username)
			data = append(data, searchUser[i].Username)
		}
	}

	searchPage := struct {
		Search string
		Media  string
		Sort   string
		Time   string
		Data   []string
	}{
		Search: search,
		Media:  media,
		Sort:   sort,
		Time:   time,
		Data:   data,
	}

	err = tmpl.Execute(w, searchPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
