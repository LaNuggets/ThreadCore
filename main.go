package main

import (
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
	db, err := sql.Open("sqlite3", "threadcore.db")
	database.CheckErr(err)

	// At the end of the program close the connnection

	query, _ := db.Prepare("INSERT INTO user (id, profilePicture, email, username, password) VALUES (?, ?, ?, ?, ?)")
	query.Exec("id", "pp", "email", "username", "password")
	defer query.Close()

	defer db.Close()

	FileServer := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	http.HandleFunc("/", Home)
	http.HandleFunc("/community", Community)
	http.HandleFunc("/friends", Friends)
	http.HandleFunc("/message", Message)
	http.HandleFunc("/post", Post)
	http.HandleFunc("/profile", Profile)
	http.HandleFunc("/404", NotFound)
	http.HandleFunc("/search", Search)
	http.HandleFunc("/connection", Connection)

	fmt.Println("Server Start at:")
	fmt.Println("http://localhost" + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
