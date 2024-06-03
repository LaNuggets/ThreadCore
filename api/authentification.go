package api

import (
	"ThreadCore/database"
	"log"
	"net/http"
	"time"

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
}

func ConnectionProfile(email string, password string, w http.ResponseWriter, r *http.Request) {

	user := database.GetUserByEmail(email)
	if CheckPasswordHash(password, user.Password) {

		expiration := time.Now().Add(2 * 24 * time.Hour)
		cookieUuid := http.Cookie{Name: "Uuid", Value: user.Uuid, Path: "/", Expires: expiration}
		http.SetCookie(w, &cookieUuid)

		cookieUsername := http.Cookie{Name: "Username", Value: user.Username, Path: "/", Expires: expiration}
		http.SetCookie(w, &cookieUsername)

		http.Redirect(w, r, "/", http.StatusSeeOther)

		println("Welcome", user.Username)
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

	user := database.User{Id: 0, Uuid: uuid, Profile: "../static/images/users/profiles/bah.png", Banner: "../static/images/users/banners/fleur.jpg", Email: email, Username: username, Password: hashedPassword}
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

// UUID
func GetNewUuid() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", uuid)
	return uuid.String()
}

// Cookies

func SetCookie(key string, value string, w http.ResponseWriter) {
	expiration := time.Now().Add(15 * 24 * time.Hour)
	cookie := http.Cookie{Name: key, Value: value, Path: "/", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func GetCookie(wantedCookie string, r *http.Request) string {
	var cookieUser *http.Cookie
	var errUser error

	cookieUser, errUser = r.Cookie(wantedCookie)
	if errUser != nil {
		if errUser == http.ErrNoCookie {
			// No cookie = Not connected
			return ""
		}
	}
	return cookieUser.Value
}

func DeleteCookie(key string, w http.ResponseWriter) {
	cookie := http.Cookie{Name: key, Value: "", Path: "/", Expires: time.Unix(0, 0)}
	http.SetCookie(w, &cookie)
}
