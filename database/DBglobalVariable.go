package database

import (
	"fmt"
	"net/http"
)

func CheckErr(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}

func ContainsPost(postList []PostInfo, postToFind PostInfo) bool {
	for _, post := range postList {
		if post.Uuid == postToFind.Uuid {
			return true
		}
	}
	return false
}

func ContainsUser(userList []User, userToFind User) bool {
	for _, user := range userList {
		if user.Uuid == userToFind.Uuid {
			return true
		}
	}
	return false
}
