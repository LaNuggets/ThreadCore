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

// !The Community function is used to create the commuity page. This page is used to display the communities based on the curent filter. She take as argument a writer and a request.
func Community(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/community/" {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/community.html")
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	communityName := strings.ReplaceAll(r.URL.Path, "/community/", "")
	if strings.Contains(communityName, "/") {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
	}

	//!ProfilePicture and connection check
	userUuid := api.GetCookie("uuid", r)
	user := database.GetUserByUuid(userUuid, w, r)

	community := database.GetCommunityByName(communityName, w, r)

	//!Following info
	followcount := len(database.GetUsersByCommunity(community.Id, w, r))
	isFollowing := database.ExistsUserCommunity(user.Id, community.Id, w, r)

	sort := r.URL.Query().Get("sort")
	ChoosenTime := r.URL.Query().Get("time")

	var sortedPosts []database.PostInfo
	var difference time.Duration

	/*
	  !Switch case for sorting communities by time(past year, past month, past week, past day, past hour) and by most post, most members.
	*/
	switch sort {
	case "popular":
		searchedPost := database.GetPostByPopularByCommunity(community.Id, w, r)
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
		searchedPost := database.GetPostsByCommunity(community.Id, w, r)
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
		searchedPost := database.GetPostByMostCommentByCommunity(community.Id, w, r)
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

	communityPage := struct {
		User            UserPageInfo
		Community       database.CommunityInfo
		CommunityExists bool
		FollowCount     int
		IsFollowing     bool
		SortedPosts     []SortedPostsInfo
	}{
		User:            userInfo,
		Community:       community,
		CommunityExists: community != database.CommunityInfo{},
		FollowCount:     followcount,
		IsFollowing:     isFollowing,
		SortedPosts:     sortedPostsInfo,
	}

	err = tmpl.Execute(w, communityPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
