package api

import (
	"ThreadCore/database"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// BCRYPT PASSWORD

/*
! HashPassword uses the x/crypto/bycrypt package to hash a string form passsword into a hashed password which will be stored in the database
*/
func HashPassword(password string, w http.ResponseWriter, r *http.Request) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	database.CheckErr(err, w, r)
	return string(bytes)
}

//BCRYPT
/*
! CheckPasswordHash uses the x/crypto/bycrypt package to take an already hashed password and a string form password to compare them and determine if the string form password is correct
*/
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//UUID
/*
! GetNewUuid uses the gofrs/uuid package to generate a random uuid to be assigned to the new users and posts created and stored in the database
*/
func GetNewUuid() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return uuid.String()
}

// Cookies
/*
! SetCookie takes a key and a value and creates a cookie for the user's browser to store for 2 weeks
? We only store the user's Uuid in this case
*/
func SetCookie(key string, value string, w http.ResponseWriter) {
	expiration := time.Now().Add(15 * 24 * time.Hour)
	cookie := http.Cookie{Name: key, Value: value, Path: "/", Expires: expiration}
	http.SetCookie(w, &cookie)
}

/*
! GetCookie returns the value of the cookie we are looking for in the user's browser
*/
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

/*
! DeleteCookie deletes the cookie we want in the user's browser
? this is used we the user disconects or deletes their account
*/
func DeleteCookie(key string, w http.ResponseWriter) {
	cookie := http.Cookie{Name: key, Value: "", Path: "/", Expires: time.Unix(0, 0)}
	http.SetCookie(w, &cookie)
}
