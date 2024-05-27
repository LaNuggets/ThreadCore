package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id             string
	profilePicture string
	email          string
	username       string
	password       string
}

func addUser(db *sql.DB, user user) {
	id := "" // ADD UUID
	query, _ := db.Prepare("INSERT INTO user (id, profilePicture, email, username, password) VALUES (?, ?, ?, ?, ?)")
	query.Exec(id, user.profilePicture, user.email, user.username, user.password)
	defer query.Close()
}

func getUserByEmail(db *sql.DB, email string) []user {
	rows, err := db.Query("SELECT * FROM user WHERE email='" + email + "'")
	defer rows.Close()

	err = rows.Err()
	checkErr(err)

	userList := make([]user, 0)

	for rows.Next() {
		user := user{}
		err = rows.Scan(&user.id, &user.profilePicture, &user.email, &user.username, &user.password)
		checkErr(err)

		userList = append(userList, user)
	}

	err = rows.Err()
	checkErr(err)

	if len(userList) > 1 {
		log.Fatal("Error : Found more than 1 user with this email")
	}

	return userList
}

func getUserById(db *sql.DB, id string) user {
	rows, _ := db.Query("SELECT * FROM user WHERE id = '" + id + "'")
	defer rows.Close()

	user := user{}

	for rows.Next() {
		rows.Scan(&user.id, &user.profilePicture, &user.email, &user.username, &user.password)
	}

	return user
}

func updateUserInfo(db *sql.DB, user user) {
	stmt, err := db.Prepare("UPDATE user set profilePicture = ?, username = ?, email = ?, password = ? where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(user.profilePicture, user.username, user.email, user.password, user.id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was affected")
	}
}

func deleteUser(db *sql.DB, userID string) {
	stmt, err := db.Prepare("DELETE FROM user where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(userID)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was deleted")
	}
}
