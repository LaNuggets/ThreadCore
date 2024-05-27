package database

import (
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

func AddUser(user User) {
	id := "" // ADD UUID
	query, _ := DB.Prepare("INSERT INTO user (id, profilePicture, email, username, password) VALUES (?, ?, ?, ?, ?)")
	query.Exec(id, user.profilePicture, user.email, user.username, user.password)
	defer query.Close()
}

func GetUserByEmail(email string) []User {
	rows, err := DB.Query("SELECT * FROM user WHERE email='" + email + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	userList := make([]User, 0)

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.id, &user.profilePicture, &user.email, &user.username, &user.password)
		CheckErr(err)

		userList = append(userList, user)
	}

	err = rows.Err()
	CheckErr(err)

	if len(userList) > 1 {
		log.Fatal("Error : Found more than 1 user with this email")
	}

	return userList
}

func GetUserById(id string) User {
	rows, _ := DB.Query("SELECT * FROM user WHERE id = '" + id + "'")
	defer rows.Close()

	user := User{}

	for rows.Next() {
		rows.Scan(&user.id, &user.profilePicture, &user.email, &user.username, &user.password)
	}

	return user
}

func UpdateUserInfo(user User) {
	stmt, err := DB.Prepare("UPDATE user set profilePicture = ?, username = ?, email = ?, password = ? where id = ?")
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(user.profilePicture, user.username, user.email, user.password, user.id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was affected")
	}
}

func DeleteUser(userID string) {
	stmt, err := DB.Prepare("DELETE FROM user where id = ?")
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(userID)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was deleted")
	}
}
