package api

import (
	"ThreadCore/database"
	"fmt"
	"net/http"
	"strconv"
)

func LikeDislike(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Check if user connected
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	rating := r.FormValue("rating")
	postid := r.FormValue("postId")
	postId, _ := strconv.Atoi(postid)
	commentid := r.FormValue("commentId")
	commentId, _ := strconv.Atoi(commentid)

	like := database.GetLikeByUserComment(user.Id, commentId)
	if (like == database.Like{}) {
		like = database.Like{Id: 0, Rating: rating, Comment_id: commentId, Post_id: postId, User_id: user.Id}
		database.AddLike(like)
	} else {
		like.Rating = rating
		database.UpdateLike(like)
	}
}
