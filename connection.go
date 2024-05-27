package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Connection(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/connection" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	errorMessage := ""
	if r.URL.Query().Get("error") == "username_taken" {
		errorMessage = "This username is already used. Choose an other please."
	}

	tmpl, err := template.ParseFiles("./templates/connection.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	// username, email, password := getIdentifier(r)
	// creationProfile(email, password, username, db ,w, r)

	data := struct {
		ErrorMessage string
	}{
		ErrorMessage: errorMessage,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func getIdentifier(r *http.Request) (*string, string, string) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	username := r.FormValue("username")

	if username == "" {
		return nil, email, password
	} else {
		return &username, email, password
	}
}

func creationProfile(email, password, username, db *sql.DB, w http.ResponseWriter, r *http.Request) {

	checkUsername := `SELECT COUNT(*) as count FROM user WHERE username = ?`
	checkEmail := `SELECT COUNT(*) as count FROM user WHERE email = ?`

	var countEmail int
	var countUsername int
	errUsername := db.QueryRow(checkUsername, username).Scan(&countUsername)
	if errUsername != nil {
		log.Fatal("Error cheking database", errUsername)
	}
	errEmail := db.QueryRow(checkEmail, email).Scan(&countEmail)
	if errEmail != nil {
		log.Fatal("Error cheking database", errEmail)
	}

	if countUsername == 1 {
		http.Redirect(w, r, "/connection?error=username_taken", http.StatusFound)
	} else if countEmail == 1 {
		http.Redirect(w, r, "/connection?error=email_taken", http.StatusFound)
	} else if countUsername == 0 && countEmail == 0 {
		_, err2 := db.Exec("INSERT INTO user(email, username, password) VALUES (?, ?, ?)", email, username, password)
		if err2 != nil {
			http.Error(w, "Error inserting data into database", http.StatusInternalServerError)
			return
		}
		fmt.Println("New user added !")
	}
}

func connectionProfile(email, password, db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
