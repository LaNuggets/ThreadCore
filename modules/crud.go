package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

db, err := sql.Open("sqlite3", "threadcore.db")
	fmt.Println("create DB:")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

func addUser() {
	
}