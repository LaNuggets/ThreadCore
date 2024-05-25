package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = ":8080"

func main() {

	FileServer := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	http.HandleFunc("/home", Home)
	http.HandleFunc("/community", Community)
	http.HandleFunc("/friends", Friends)
	http.HandleFunc("/message", Message)
	http.HandleFunc("/post", Post)
	http.HandleFunc("/profile", Profile)
	http.HandleFunc("/404", NotFound)
	http.HandleFunc("/search", Search)
	http.HandleFunc("/connection", Connection)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	})

	fmt.Println("Server Start at:")
	fmt.Println("http://localhost" + port + "/home")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
