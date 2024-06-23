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

// ! The User function is used to create the user profile page. This page display all the information about ann user like is nulber of follower, his post
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

	//ProfilePicture and connection check
	userUuid := api.GetCookie("uuid", r)
	user := database.GetUserByUuid(userUuid, w, r)

	profile := database.GetUserByUsername(username, w, r)

	//Following info
	followcount := len(database.GetFriendsByUser(profile.Id, w, r))
	isFollowing := database.ExistsFriend(user.Id, profile.Id, w, r)

	sort := r.URL.Query().Get("sort")
	ChoosenTime := r.URL.Query().Get("time")

	var sortedPosts []database.PostInfo
	var difference time.Duration
	/*
	  !Switch case for sorting post by time(past year, past month, past week, past day, past hour) and by most popular, newest and most comment.
	*/
	switch sort {
	case "popular":
		searchedPost := database.GetPostByPopularByUser(profile.Id, w, r)
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
		searchedPost := database.GetPostsByUser(profile.Id, w, r)
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
		searchedPost := database.GetPostByMostCommentByUser(profile.Id, w, r)
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

	type SortedPostsInfo struct {
		Post         database.PostInfo
		PostLikes    int
		PostDislikes int
		UserRating   string
	}

	var sortedPostsInfo []SortedPostsInfo
	//!Get rating info on each posts
	for i := 0; i < len(sortedPosts); i++ {
		likes := 0
		dislikes := 0
		allRatings := database.GetLikesByPost(sortedPosts[i].Id, w, r)
		for i := 0; i < len(allRatings); i++ {
			if allRatings[i].Rating == "like" {
				likes += 1
			} else if allRatings[i].Rating == "dislike" {
				dislikes += 1
			}
		}
		userRating := database.GetLikeByUserPost(user.Id, sortedPosts[i].Id, w, r).Rating
		sortedPost := SortedPostsInfo{Post: sortedPosts[i], PostLikes: likes, PostDislikes: dislikes, UserRating: userRating}
		sortedPostsInfo = append(sortedPostsInfo, sortedPost)
	}

	type UserPageInfo struct {
		Connected bool
		Profile   string
		Username  string
		Uuid      string
	}
	userInfo := UserPageInfo{
		Connected: userUuid != "",
		Profile:   user.Profile,
		Username:  user.Username,
		Uuid:      user.Uuid,
	}

	type ProfilePageInfo struct {
		Uuid     string
		Profile  string
		Banner   string
		Username string
		Email    string
	}
	profileExists := profile != database.User{}
	profileInfo := ProfilePageInfo{}

	if profileExists {
		profileInfo = ProfilePageInfo{
			Uuid:     user.Uuid,
			Profile:  profile.Profile,
			Banner:   profile.Banner,
			Username: profile.Username,
			Email:    profile.Email,
		}
	}

	userPage := struct {
		User          UserPageInfo
		Profile       ProfilePageInfo
		ProfileExists bool
		FollowCount   int
		IsFollowing   bool
		SortedPosts   []SortedPostsInfo
	}{
		User:          userInfo,
		Profile:       profileInfo,
		ProfileExists: profileExists,
		FollowCount:   followcount,
		IsFollowing:   isFollowing,
		SortedPosts:   sortedPostsInfo,
	}

	err = tmpl.Execute(w, userPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
