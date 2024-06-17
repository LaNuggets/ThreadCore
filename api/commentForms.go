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
		http.Redirect(w, r, "/search/?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/search/?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	}

	commentid := r.FormValue("commentId")
	commentId, _ := strconv.Atoi(commentid)
	postid := r.FormValue("postId")
	postId, _ := strconv.Atoi(postid)
	content := r.FormValue("content")

	comment := database.Comment{Id: 0, User_id: user.Id, Post_id: postId, Comment_id: commentId, Content: content, Created: time.Now()}
	database.AddComment(comment)

	postUuid := r.FormValue("postUuid")
	http.Redirect(w, r, "/post/"+postUuid+"?type=success&message=Comment+posted+!", http.StatusSeeOther)
}

// UPDATE EXISTING comment
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	commentId := r.FormValue("commentId")
	commentid, _ := strconv.Atoi(commentId)
	comment := database.GetCommentById(commentid, w, r)
	if (comment == database.CommentInfo{}) {
		fmt.Println("comment does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/search/?type=error&message=Comment+not+found+!", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	} else if comment.User_id != user.Id {
		fmt.Println("user not author of comment") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=User+not+alowed+to+do+this+action+!", http.StatusSeeOther)
		return
	}

	postid := r.FormValue("postId")
	postId, _ := strconv.Atoi(postid)
	content := r.FormValue("content")

	commentUpdate := database.Comment{Id: 0, User_id: user.Id, Post_id: postId, Comment_id: commentid, Content: content, Created: time.Now()}
	database.AddComment(commentUpdate)

	http.Redirect(w, r, "/comment/"+commentId+"?type=success&message=Comment+successfully+update+!", http.StatusSeeOther)
}

// DELETE comment
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	commentId := r.FormValue("commentId")
	commentid, _ := strconv.Atoi(commentId)
	comment := database.GetCommentById(commentid, w, r)
	if (comment == database.CommentInfo{}) {
		fmt.Println("comment does not exist") // TO-DO : send error comment not found
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=Comment+not+found+!", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	} else if comment.User_id != user.Id {
		fmt.Println("user not author of comment") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/comment/?type=error&message=User+not+alowed+to+do+this+action+!"+commentId, http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/comment/"+commentId+"?type=error&message=Confirm+deletion+!", http.StatusSeeOther)
		return
	} else {
		database.DeleteComment(comment.Id, w, r)
	}

	//Send confirmation message
	http.Redirect(w, r, "/?type=success&message=Comment+deleted+!", http.StatusSeeOther)
}
