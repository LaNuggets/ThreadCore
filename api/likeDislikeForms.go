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
		http.Redirect(w, r, "/search/?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/search/?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	}

	rating := r.FormValue("rating")
	postid := r.FormValue("postId")
	postId, _ := strconv.Atoi(postid)
	commentid := r.FormValue("commentId")
	commentId, _ := strconv.Atoi(commentid)

	id := r.FormValue("postUuid")
	Id, _ := strconv.Atoi(id)
	post := database.GetPostById(Id, w, r)
	postUuid := post.Uuid

	if post.User_id == user.Id {
		fmt.Println("cannot dislike your own post") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/post/"+postUuid+"?type=error&message=You+cannot+dislike+your+own+post+!", http.StatusSeeOther)
		return
	}

	var like database.Like
	if commentId == 0 {
		like = database.GetLikeByUserPost(user.Id, postId, w, r)
	} else {
		like = database.GetLikeByUserComment(user.Id, commentId, w, r)
	}

	if (like == database.Like{}) {
		like = database.Like{Id: 0, Rating: rating, Comment_id: commentId, Post_id: postId, User_id: user.Id}
		database.AddLike(like, w, r)
	} else {
		if like.Rating == rating {
			database.DeleteLike(like.Id, w, r)
		} else {
			like.Rating = rating
			database.UpdateLike(like, w, r)
		}
	}

	http.Redirect(w, r, "/post/"+postUuid, http.StatusSeeOther)
}
