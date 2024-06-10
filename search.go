package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
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
	media := r.URL.Query().Get("media") // media options  : posts, communities, users
	sort := r.URL.Query().Get("sort")
	ChoosenTime := r.URL.Query().Get("time")

	var sortedPosts []database.PostInfo
	var sortedCommunities []database.CommunityDisplay
	var sortedUsers []database.User

	switch media {
	case "posts":
		switch sort {
		case "popular":
			searchedPost := database.GetPostByPopular(search)
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		case "new":
			searchedPost := database.GetPostsBySearchString(search)
			api.NewestPost(searchedPost)
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		case "most_comments":
			searchedPost := database.GetPostByMostComment(search)
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var monthTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(monthTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var weekTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(weekTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var dayTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(dayTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var hourTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(hourTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		}

	case "communities":
		switch sort {
		case "popular":
			sortedCommunities = database.GetCommunitiesByNMembers(search)
		case "new":
			sortedCommunities = database.GetCommunitiesByMostPost(search)
		}

	case "users":
		switch sort {
		case "popular":
			sortedUsers = database.GetUserByMostPopular(search)
		case "new":
			sortedUsers = database.GetUserByMostPost(search)
		case "most_comments":
			sortedUsers = database.GetUserByMostComment(search)
		}
	}

	//! just for testing
	for i := 0; i < len(sortedPosts); i++ {
		fmt.Println(sortedPosts[i].Title)
	}
	for i := 0; i < len(sortedCommunities); i++ {
		fmt.Println(sortedCommunities[i].Name)
	}
	for i := 0; i < len(sortedUsers); i++ {
		fmt.Println(sortedUsers[i].Username)
	}
	//! just for testing

	searchPage := struct {
		Connected bool
		// Search            string
		// Media             string
		// Sort              string
		// Time              string
		// SortedPosts       []database.PostInfo
		// SortedCommunities []database.CommunityDisplay
		// SortedUsers       []database.User
	}{
		Connected: api.GetCookie("uuid", r) != "",
		// Search:            search,
		// Media:             media,
		// Sort:              sort,
		// Time:              ChoosenTime,
		// SortedPosts:       sortedPosts,
		// SortedCommunities: sortedCommunities,
		// SortedUsers:       sortedUsers,
	}

	err = tmpl.Execute(w, searchPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
