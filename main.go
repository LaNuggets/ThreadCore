package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var port = ":8080"

func main() {

	// Open the database connection in the global varaible DB located in database/DBglobalVariable.go
	var err error
	database.DB, err = sql.Open("sqlite3", "threadcore.db")
	database.CheckErr(err)

	// At the end of the program close the connnection
	defer database.DB.Close()

	FileServer := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	http.HandleFunc("/", Home)
	http.HandleFunc("/community/", Community)
	http.HandleFunc("/friends", Friends)
	http.HandleFunc("/message", Message)
	http.HandleFunc("/post/", Post)
	http.HandleFunc("/user/", User)
	http.HandleFunc("/404", NotFound)
	http.HandleFunc("/search/", Search)
	http.HandleFunc("/connection", Connection)

	// Forms routes
	http.HandleFunc("/createCommunity", api.CreateCommunity)
	http.HandleFunc("/updateCommunity", api.UpdateCommunity)
	// http.HandleFunc("/deleteCommunity", DeleteCommunity)
	// http.HandleFunc("/createPost", CreatePost)
	// http.HandleFunc("/updatePost", UpdatePost)
	// http.HandleFunc("/deletePost", DeletePost)
	// http.HandleFunc("/createComment", CreateComment)
	// http.HandleFunc("/updateComment", UpdateComment)
	// http.HandleFunc("/deleteComment", DeleteComment)
	// http.HandleFunc("/likeDislike", LikeDislike)

	fmt.Println("Server Start at:")
	fmt.Println("http://localhost" + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
