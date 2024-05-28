package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id             string
	ProfilePicture string
	Email          string
	Username       string
	Password       string
}

func AddUser(user User) {
	id := "" // ADD UUID
	query, _ := DB.Prepare("INSERT INTO user (id, profilePicture, email, username, password) VALUES (?, ?, ?, ?, ?)")
	query.Exec(id, user.ProfilePicture, user.Email, user.Username, user.Password)
	defer query.Close()
}

func GetUserByEmail(email string) User {
	rows, err := DB.Query("SELECT * FROM user WHERE email='" + email + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.ProfilePicture, &user.Email, &user.Username, &user.Password)
	}

	return user
}

func GetUserById(id string) User {
	rows, _ := DB.Query("SELECT * FROM user WHERE id = '" + id + "'")
	defer rows.Close()

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.ProfilePicture, &user.Email, &user.Username, &user.Password)
	}

	return user
}

func UpdateUserInfo(user User) {
	stmt, err := DB.Prepare("UPDATE user set profilePicture = ?, username = ?, email = ?, password = ? where id = ?")
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(user.ProfilePicture, user.Username, user.Email, user.Password, user.Id)
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
