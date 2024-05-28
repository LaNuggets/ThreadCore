package api

import (
	"ThreadCore/database"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Authentication(w http.ResponseWriter, r *http.Request) {
	username, email, password := GetIdentifier(r)
	if email != "" {
		ChooseConnectionOrCreation(username, email, password, w, r)
	}
}

func GetIdentifier(r *http.Request) (*string, string, string) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username == "" {
		return nil, email, password
	} else {
		return &username, email, password
	}
}

func ChooseConnectionOrCreation(username *string, email string, password string, w http.ResponseWriter, r *http.Request) {
	if username == nil {
		ConnectionProfile(email, password, w, r)
	} else {
		CreationProfile(*username, email, password)
	}
}

func CreationProfile(username string, email string, password string) {
	user := AddUserValue(username, email, password)
	database.AddUser(user)
	println(user.Username)
	println(user.Email)
	println(user.Password)
	println("Creation successful")
}

func ConnectionProfile(email string, password string, w http.ResponseWriter, r *http.Request) {
	user := database.GetUserByEmail(email)
	if user.Password == password {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		println("Welcome", user.Username)
		//Todo: Cookie
	} else {
		http.Redirect(w, r, "/connection?error=password_taken", http.StatusFound)
	}

}

func AddUserValue(username string, email string, password string) database.User {
	u, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u)
	uuid := u.String()

	hashedPassword := HashPassword(password)

	user := database.User{nil, uuid, "picture", email, username, hashedPassword}
	return user
}

// BCRYPT PASSWORD

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	database.CheckErr(err)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
