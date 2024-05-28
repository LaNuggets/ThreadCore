package api

import (
	"ThreadCore/database"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Authentication(w http.ResponseWriter, r *http.Request) (*http.Cookie, *http.Cookie) {
	var cookieUui, cookieUser *http.Cookie

	username, email, password := GetIdentifier(r)
	if email != "" {
		cookieUui, cookieUser = ChooseConnectionOrCreation(username, email, password, w, r)
	}
	return cookieUser, cookieUui
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

func ChooseConnectionOrCreation(username *string, email string, password string, w http.ResponseWriter, r *http.Request) (*http.Cookie, *http.Cookie) {
	var cookieUui, cookieUser *http.Cookie
	if username == nil {
		cookieUui, cookieUser = ConnectionProfile(email, password, w, r)
	} else {
		CreationProfile(*username, email, password)
	}
	return cookieUui, cookieUser
}

func CreationProfile(username string, email string, password string) {
	user := AddUserValue(username, email, password)
	database.AddUser(user)
	println(user.Username)
	println(user.Email)
	println(user.Password)
	println("Creation successful")
}

func ConnectionProfile(email string, password string, w http.ResponseWriter, r *http.Request) (*http.Cookie, *http.Cookie) {
	var cookieUui, cookieUser *http.Cookie

	user := database.GetUserByEmail(email)
	if CheckPasswordHash(password, user.Password) {
		http.Redirect(w, r, "/", http.StatusSeeOther)

		expiration := time.Now().Add(2 * 24 * time.Hour)
		cookieUuid := http.Cookie{Name: "Uuid", Value: user.Uuid, Path: "/", Expires: expiration}
		http.SetCookie(w, &cookieUuid)

		cookieUsername := http.Cookie{Name: "Username", Value: user.Username, Path: "/", Expires: expiration}
		http.SetCookie(w, &cookieUsername)

		var errUuid error
		var errUser error

		cookieUui, errUuid = r.Cookie("Uuid")
		if errUuid != nil {
			if errUuid == http.ErrNoCookie {
				// Si le cookie n'existe pas
				log.Fatal("Cookie uuid not found")
				// Vous pouvez gérer ce cas en définissant une valeur par défaut, en renvoyant une erreur HTTP, etc.
				log.Fatal("Cookie 'uuid' not found", http.StatusUnauthorized)
			}
			// Si une autre erreur s'est produite
			log.Fatal("Error retrieving cookie:", errUuid)
		}

		cookieUser, errUser = r.Cookie("Username")
		if errUser != nil {
			if errUser == http.ErrNoCookie {
				// Si le cookie n'existe pas
				log.Fatal("Cookie username not found")
				// Vous pouvez gérer ce cas en définissant une valeur par défaut, en renvoyant une erreur HTTP, etc.
				log.Fatal("Cookie 'uuid' not found", http.StatusUnauthorized)
			}
			// Si une autre erreur s'est produite
			log.Fatal("Error retrieving cookie:", errUser)
		}

		fmt.Fprintln(w, cookieUui, cookieUser)

		println("Welcome", user.Username)
	} else {
		http.Redirect(w, r, "/connection?error=password_taken", http.StatusFound)
	}
	return cookieUui, cookieUser
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
