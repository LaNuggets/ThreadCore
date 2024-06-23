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

// ! The Post function is used to create the post page. This page diplay all the information about a certin post like the content, is creator, the comment, the number of like/dislike.
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

	url := strings.ReplaceAll(r.URL.Path, "/post/", "")
	urlQuery := strings.Split(url, "?")
	postUuid := urlQuery[0]

	if strings.Contains(postUuid, "/") {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
	}

	post := database.GetPostByUuid(postUuid, w, r)
	comments := database.GetCommentsByPost(post.Id, w, r)
	community := database.GetCommunityById(post.Community_id, w, r)

	//!Time formating for the post
	difference := time.Now().Sub(post.Created)
	post.Time = api.GetFormatedDuration(difference)

	type CommentsPageInfo struct {
		Comment         database.CommentInfo
		CommentLikes    int
		CommentDislikes int
		UserRating      string
		SubComments     []database.CommentInfo
	}

	//!ProfilePicture and connection check
	userUuid := api.GetCookie("uuid", r)
	user := database.GetUserByUuid(userUuid, w, r)
	// Delete cookie if not existing in database
	if (user == database.User{} && userUuid != "") {
		api.DeleteCookie(userUuid, w)
	}

	commentInfo := []CommentsPageInfo{}

	//!Time formating for comments + get comment answers
	for i := 0; i < len(comments); i++ {
		difference := time.Now().Sub(comments[i].Created)
		comments[i].Time = api.GetFormatedDuration(difference)
		subComments := database.GetCommentsByComment(comments[i].Id, w, r)
		for i := 0; i < len(subComments); i++ {
			difference := time.Now().Sub(subComments[i].Created)
			subComments[i].Time = api.GetFormatedDuration(difference)
		}
		likes := 0
		dislikes := 0
		allRatings := database.GetLikesByComment(comments[i].Id, w, r)
		for i := 0; i < len(allRatings); i++ {
			if allRatings[i].Rating == "like" {
				likes += 1
			} else if allRatings[i].Rating == "dislike" {
				dislikes += 1
			}
		}
		userRating := database.GetLikeByUserComment(user.Id, comments[i].Id, w, r).Rating
		fullComment := CommentsPageInfo{Comment: comments[i], CommentLikes: likes, CommentDislikes: dislikes, UserRating: userRating, SubComments: subComments}
		commentInfo = append(commentInfo, fullComment)
	}
	//!Check if the post is in a community
	followcount := 0
	var isFollowing bool
	if post.Community_id == 0 {
		followcount = len(database.GetFriendsByUser(post.User_id, w, r))
		isFollowing = database.ExistsFriend(post.User_id, user.Id, w, r)
	} else {
		followcount = len(database.GetUsersByCommunity(community.Id, w, r))
		isFollowing = database.ExistsUserCommunity(user.Id, post.Community_id, w, r)
	}

	// !Get likes and dislikes
	likes := 0
	dislikes := 0
	allRatings := database.GetLikesByPost(post.Id, w, r)
	for i := 0; i < len(allRatings); i++ {
		if allRatings[i].Rating == "like" {
			likes += 1
		} else if allRatings[i].Rating == "dislike" {
			dislikes += 1
		}
	}

	userRating := database.GetLikeByUserPost(user.Id, post.Id, w, r).Rating

	type UserPageInfo struct {
		Connected bool
		Profile   string
		Username  string
		Uuid      string
		Id        int
	}
	userInfo := UserPageInfo{
		Connected: user.Uuid != "",
		Profile:   user.Profile,
		Username:  api.GetCookie("username", r),
		Uuid:      user.Uuid,
		Id:        user.Id,
	}

	postPage := struct {
		User         UserPageInfo
		Post         database.PostInfo
		PostExists   bool
		PostLikes    int
		PostDislikes int
		UserRating   string
		Community    database.CommunityInfo
		Comments     []CommentsPageInfo
		FollowCount  int
		IsFollowing  bool
	}{
		User:         userInfo,
		Post:         post,
		PostExists:   post != database.PostInfo{},
		PostLikes:    likes,
		PostDislikes: dislikes,
		UserRating:   userRating,
		Community:    community,
		Comments:     commentInfo,
		FollowCount:  followcount,
		IsFollowing:  isFollowing,
	}

	err = tmpl.Execute(w, postPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
