package main

import (
	"ThreadCore/api"
	"fmt"
	"log"
	"net/http"
)

var port = ":8080"

// ! The main function is where the programme start. The function initialize all the page path. This is the root of the forums.
func main() {
	FileServer := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", FileServer))
	//! All pages redirection path
	http.HandleFunc("/", Home)
	http.HandleFunc("/community/", Community)
	http.HandleFunc("/post/", Post)
	http.HandleFunc("/user/", User)
	http.HandleFunc("/404", NotFound)
	http.HandleFunc("/500", Error)
	http.HandleFunc("/search/", Search)

	// !Forms routes
	http.HandleFunc("/createCommunity", api.CreateCommunity)
	http.HandleFunc("/updateCommunity", api.UpdateCommunity)
	http.HandleFunc("/deleteCommunity", api.DeleteCommunity)
	http.HandleFunc("/createPost", api.CreatePost)
	http.HandleFunc("/updatePost", api.UpdatePost)
	http.HandleFunc("/deletePost", api.DeletePost)
	http.HandleFunc("/createComment", api.CreateComment)
	http.HandleFunc("/updateComment", api.UpdateComment)
	http.HandleFunc("/deleteComment", api.DeleteComment)
	http.HandleFunc("/likeDislike", api.LikeDislike)
	http.HandleFunc("/login", api.Login)
	http.HandleFunc("/signup", api.Signup)
	http.HandleFunc("/disconnect", api.Disconnect)
	http.HandleFunc("/updateUserInfo", api.UpdateUser)
	http.HandleFunc("/deleteUserInfo", api.DeleteUser)
	http.HandleFunc("/followCommunity", api.FollowCommunity)
	http.HandleFunc("/unfollowCommunity", api.UnfollowCommunity)
	http.HandleFunc("/followUser", api.FollowUser)
	http.HandleFunc("/unfollowUser", api.UnfollowUser)

	fmt.Println("Server Start at:")
	fmt.Println("http://localhost" + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
