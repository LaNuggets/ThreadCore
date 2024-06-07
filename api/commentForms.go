package api

import (
	"ThreadCore/database"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// CREATE NEW COMMENT
func CreateComment(w http.ResponseWriter, r *http.Request) {
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

	commentid := r.FormValue("commentId")
	commentId, _ := strconv.Atoi(commentid)
	postid := r.FormValue("postId")
	postId, _ := strconv.Atoi(postid)
	content := r.FormValue("content")

	comment := database.Comment{Id: 0, User_id: user.Id, Post_id: postId, Comment_id: commentId, Content: content, Created: time.Now()}
	database.AddComment(comment)

	//http.Redirect(w, r, "/comment/"+name, http.StatusSeeOther)
}

// UPDATE EXISTING comment
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	commentId := r.FormValue("commentId")
	commentid, _ := strconv.Atoi(commentId)
	comment := database.GetCommentById(commentid)
	if (comment == database.CommentInfo{}) {
		fmt.Println("comment does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
		return
	} else if comment.User_id != user.Id {
		fmt.Println("user not author of comment") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
		return
	}

	postid := r.FormValue("postId")
	postId, _ := strconv.Atoi(postid)
	content := r.FormValue("content")

	commentUpdate := database.Comment{Id: 0, User_id: user.Id, Post_id: postId, Comment_id: commentid, Content: content, Created: time.Now()}
	database.AddComment(commentUpdate)

	http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
}

// DELETE comment
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	commentId := r.FormValue("commentId")
	commentid, _ := strconv.Atoi(commentId)
	comment := database.GetCommentById(commentid)
	if (comment == database.CommentInfo{}) {
		fmt.Println("comment does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
		return
	} else if comment.User_id != user.Id {
		fmt.Println("user not author of comment") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/comment/"+commentId, http.StatusSeeOther)
		return
	} else {
		database.DeleteComment(comment.Id)
	}

	//Send confirmation message
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
