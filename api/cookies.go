package api

import "net/http"

func CookieGetter(wentedCookie string, r *http.Request) string {
	var cookieUser *http.Cookie
	var errUser error

	cookieUser, errUser = r.Cookie(wentedCookie)
	if errUser != nil {
		if errUser == http.ErrNoCookie {
			// No cookie = Not connected
			return ""
		}
	}
	return cookieUser.Value
}