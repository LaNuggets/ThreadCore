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

	userUuid := api.GetCookie("uuid", r)
	user := database.GetUserByUuid(userUuid, w, r)
	userProfile := user.Profile
	username := user.Username

	search := r.URL.Query().Get("q")
	media := r.URL.Query().Get("media") // media options  : posts, communities, users
	sort := r.URL.Query().Get("sort")
	ChoosenTime := r.URL.Query().Get("time")

	var sortedPosts []database.PostInfo
	var sortedCommunities []database.CommunityDisplay
	var sortedUsers []database.User
	var difference time.Duration

	switch media {
	case "posts":
		switch sort {
		case "popular":
			searchedPost := database.GetPostByPopular(search, w, r)
			switch ChoosenTime {
			case "all_time":
				for i := 0; i < len(searchedPost); i++ {
					difference = time.Now().Sub(searchedPost[i].Created)
					searchedPost[i].Time = api.GetFormatedDuration(difference)
				}
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		case "new":
			searchedPost := database.GetPostsBySearchString(search, w, r)
			api.NewestPost(searchedPost)
			switch ChoosenTime {
			case "all_time":
				for i := 0; i < len(searchedPost); i++ {
					difference = time.Now().Sub(searchedPost[i].Created)
					searchedPost[i].Time = api.GetFormatedDuration(difference)
				}
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		case "most_comments":
			searchedPost := database.GetPostByMostComment(search, w, r)
			switch ChoosenTime {
			case "all_time":
				for i := 0; i < len(searchedPost); i++ {
					difference = time.Now().Sub(searchedPost[i].Created)
					searchedPost[i].Time = api.GetFormatedDuration(difference)
				}
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var monthTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(monthTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var weekTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(weekTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var dayTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(dayTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var hourTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(hourTime)) {
						//Time formating for the post
						difference = time.Now().Sub(searchedPost[i].Created)
						searchedPost[i].Time = api.GetFormatedDuration(difference)
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		}

	case "communities":
		switch sort {
		case "popular":
			sortedCommunities = database.GetCommunitiesByNMembers(search, w, r)
		case "new":
			sortedCommunities = database.GetCommunitiesByMostPost(search, w, r)
		}

	case "users":
		switch sort {
		case "popular":
			sortedUsers = database.GetUserByMostPopular(search, w, r)
		case "new":
			sortedUsers = database.GetUserByMostPost(search, w, r)
		case "most_comments":
			sortedUsers = database.GetUserByMostComment(search, w, r)
		}
	}

	//! just for testing
	for i := 0; i < len(sortedPosts); i++ {
		fmt.Println(sortedPosts[i].Time)
	}
	for i := 0; i < len(sortedCommunities); i++ {
		fmt.Println(sortedCommunities[i].Name)
	}
	for i := 0; i < len(sortedUsers); i++ {
		fmt.Println(sortedUsers[i].Username)
	}
	//! just for testing

	searchPage := struct {
		Connected         bool
		Profile           string
		Search            string
		Media             string
		Sort              string
		Time              string
		Username          string
		SortedPosts       []database.PostInfo
		SortedCommunities []database.CommunityDisplay
		SortedUsers       []database.User
	}{
		Connected:         userUuid != "",
		Profile:           userProfile,
		Search:            search,
		Media:             media,
		Sort:              sort,
		Time:              ChoosenTime,
		Username:          username,
		SortedPosts:       sortedPosts,
		SortedCommunities: sortedCommunities,
		SortedUsers:       sortedUsers,
	}

	err = tmpl.Execute(w, searchPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
