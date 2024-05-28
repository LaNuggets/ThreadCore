package database

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id             *int
	Uuid           string
	ProfilePicture string
	Email          string
	Username       string
	Password       string
}

func AddUser(user User) {
	query, _ := DB.Prepare("INSERT INTO user (uuid, profilePicture, email, username, password) VALUES (?, ?, ?, ?, ?)")
	query.Exec(user.Uuid, user.ProfilePicture, user.Email, user.Username, user.Password)
	defer query.Close()
	fmt.Println("test")
}

func GetUserByEmail(email string) User {
	rows, err := DB.Query("SELECT * FROM user WHERE email='" + email + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.ProfilePicture, &user.Email, &user.Username, &user.Password)
	}

	return user
}

func GetUserById(id string) User {
	rows, _ := DB.Query("SELECT * FROM user WHERE id = '" + id + "'")
	defer rows.Close()

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.ProfilePicture, &user.Email, &user.Username, &user.Password)
	}

	return user
}

func UpdateUserInfo(user User) {
	query, err := DB.Prepare("UPDATE user set uuid = ?, profilePicture = ?, username = ?, email = ?, password = ? where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(user.Uuid, user.ProfilePicture, user.Username, user.Email, user.Password, user.Id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was affected")
	}
}

func DeleteUser(userID string) {
	query, err := DB.Prepare("DELETE FROM user where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(userID)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was deleted")
	}
}
