package database

import (
	"fmt"
	"net/http"
)

func CheckErr(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}
