package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var port = ":8080"

func main() {
	FileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	http.HandleFunc("/", home)

	fmt.Println("Server Start at:")
	fmt.Println("http://localhost" + port)

	error := http.ListenAndServe(port, nil)

	if errors.Is(error, http.ErrServerClosed) {
		fmt.Println("Server Close")
	} else if error != nil {
		fmt.Println("Error starting server: s%", error)
		os.Exit(1)
	}
}

func home(w http.ResponseWriter, r *http.Request) { // Lunch a new page for the lose condition
	tmpl := template.Must(template.ParseFiles("./templates/home.html")) // Read the home page

	tmpl.Execute(w, nil)
}
