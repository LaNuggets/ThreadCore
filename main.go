package main

import (
	"ThreadCore/database"
	"database/sql"
	"fmt"
	"log"
)

func main() {
	var err error
	database.DB, err = sql.Open("sqlite3", "threadcore.db")
	fmt.Println("create DB:")
	if err != nil {
		log.Fatal(err)
	}

	defer database.DB.Close()
}
