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
