package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
	"time"
)

// ! The Search function is used to create the search page. This page allow user find search poost, communities and other user by using filter and search bar. She take as argument a writer and a request.
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

	/*
		!Get user with cookies
	*/
	userUuid := api.GetCookie("uuid", r)
	user := database.GetUserByUuid(userUuid, w, r)
	// Delete cookie if not existing in database
	if (user == database.User{} && userUuid != "") {
		api.DeleteCookie(userUuid, w)
	}

	search := r.URL.Query().Get("q")
	media := r.URL.Query().Get("media") // media options  : posts, communities, users
	sort := r.URL.Query().Get("sort")
	ChoosenTime := r.URL.Query().Get("time")

	var sortedPosts []database.PostInfo
	var sortedCommunities []database.CommunityInfo
	var sortedUsers []database.User
	var difference time.Duration

	/*
	  !Switch case for sorting post, communities, user by time(past year, past month, past week, past day, past hour) and by most popular, newest, most comment, most post, most members
	*/

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
						//Time formating for the posts
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
			sortedCommunities = database.GetCommunitiesByMostMembers(search, w, r)
		case "most_comments":
			sortedCommunities = database.GetCommunitiesByMostPost(search, w, r)
		}

	case "users":
		switch sort {
		case "popular":
			sortedUsers = database.GetUserByMostPopular(search, w, r)
		case "most_comments":
			sortedUsers = database.GetUserByMostPost(search, w, r)
		}
	}

	/*
		!Structur to send data to html
	*/

	type UserPageInfo struct {
		Connected bool
		Profile   string
		Username  string
		Uuid      string
	}
	userInfo := UserPageInfo{
		Connected: user.Uuid != "",
		Profile:   user.Profile,
		Username:  user.Username,
		Uuid:      user.Uuid,
	}

	type ProfileInfo struct {
		Uuid     string
		Profile  string
		Banner   string
		Username string
	}

	searchPage := struct {
		User              UserPageInfo
		SortedPosts       []database.PostInfo
		SortedCommunities []database.CommunityInfo
		SortedUsers       []database.User
	}{
		User:              userInfo,
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
