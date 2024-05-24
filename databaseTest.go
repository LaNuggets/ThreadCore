package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to database
	db, err := sql.Open("sqlite3", "./ThreadCore.db")
	checkErr(err)
	// defer close
	defer db.Close()

	dropTables := "DROP TABLE IF EXISTS user; DROP TABLE IF EXISTS message; DROP TABLE IF EXISTS post; DROP TABLE IF EXISTS comment; DROP TABLE IF EXISTS like; DROP TABLE IF EXISTS friend; DROP TABLE IF EXISTS community; DROP TABLE IF EXISTS user_community; DROP TABLE IF EXISTS group; DROP TABLE IF EXISTS user_group"
	_, err = db.Exec(dropTables)

	createTables := "CREATE TABLE cars(id INTEGER PRIMARY KEY, profilePicture TEXT, price INT)"
	_, err = db.Exec(createTables)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
