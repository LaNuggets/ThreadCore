package database

import (
	"database/sql"
	"fmt"
	"net/http"
)

var DB *sql.DB

func CheckErr(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/500", http.StatusSeeOther)
	}
}
