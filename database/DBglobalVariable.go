package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
