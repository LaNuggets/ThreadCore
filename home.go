package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("./templates/home.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	//ProfilePicture and connection check
	userUuid := api.GetCookie("uuid", r)
	user := database.GetUserByUuid(userUuid, w, r)

	ChoosenTime := r.URL.Query().Get("time")

	var sortedPosts []database.PostInfo
	var difference time.Duration
	var searchedPost []database.PostInfo

	if (user == database.User{}) {
		sortedPosts = database.GetAllPosts(w, r)
	} else {
		//Get all users and communities that the user is following
		friendList := database.GetFriendsByUser(user.Id, w, r)
		communitiesList := database.GetCommunitiesByUser(user.Id, w, r)
		friendList = append(friendList, user)

		for i := 0; i < len(friendList); i++ {
			posts := database.GetPostByPopularByUser(friendList[i].Id, w, r)
			for i := 0; i < len(posts); i++ {
				searchedPost = append(searchedPost, posts[i])
			}
		}
		for i := 0; i < len(communitiesList); i++ {
			posts := database.GetPostByPopularByCommunity(communitiesList[i].Id, w, r)
			for i := 0; i < len(posts); i++ {
				searchedPost = append(searchedPost, posts[i])
			}
		}

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
	}

	type SortedPostsInfo struct {
		Post         database.PostInfo
		PostLikes    int
		PostDislikes int
		UserRating   string
	}

	var sortedPostsInfo []SortedPostsInfo
	//Get rating info on each posts
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

	homePage := struct {
		User        UserPageInfo
		SortedPosts []SortedPostsInfo
	}{
		User:        userInfo,
		SortedPosts: sortedPostsInfo,
	}
	err = tmpl.Execute(w, homePage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
