package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
