package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	id             string
	profilePicture string
	email          string
	username       string
	password       string
}

func AddUser(db *sql.DB, user User) {
	id := "" // ADD UUID
	query, _ := db.Prepare("INSERT INTO user (id, profilePicture, email, username, password) VALUES (?, ?, ?, ?, ?)")
	query.Exec(id, user.profilePicture, user.email, user.username, user.password)
	defer query.Close()
}

func GetUserByEmail(db *sql.DB, email string) []User {
	rows, err := db.Query("SELECT * FROM user WHERE email='" + email + "'")
	defer rows.Close()

	err = rows.Err()
	checkErr(err)

	people := make([]User, 0)

	for rows.Next() {
		ourPerson := User{}
		err = rows.Scan(&ourPerson.id, &ourPerson.profilePicture, &ourPerson.email, &ourPerson.username, &ourPerson.password)
		checkErr(err)

		people = append(people, ourPerson)
	}

	err = rows.Err()
	checkErr(err)

	if len(people) > 1 {
		log.Fatal("Error : Found more than 1 user with this email")
	}

	return people
}

func GetUserById(db *sql.DB, id string) User {
	rows, _ := db.Query("SELECT * FROM user WHERE id = '" + id + "'")
	defer rows.Close()

	ourPerson := User{}

	for rows.Next() {
		rows.Scan(&ourPerson.id, &ourPerson.profilePicture, &ourPerson.email, &ourPerson.username, &ourPerson.password)
	}

	return ourPerson
}

func UpdateUserInfo(db *sql.DB, user User) int64 {
	stmt, err := db.Prepare("UPDATE user set profilePicture = ?, username = ?, email = ?, password = ? where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(user.profilePicture, user.username, user.email, user.password, user.id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	return affected
}

func DeleteUser(db *sql.DB, userID string) int64 {
	stmt, err := db.Prepare("DELETE FROM user where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(userID)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	return affected
}
