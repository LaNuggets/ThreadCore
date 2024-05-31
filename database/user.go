package database

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Uuid     string
	Profile  string
	Banner   string
	Email    string
	Username string
	Password string
}

func AddUser(user User) {
	query, _ := DB.Prepare("INSERT INTO user (uuid, profile, banner, email, username, password) VALUES (?, ?, ?, ?, ?, ?)")
	query.Exec(user.Uuid, user.Profile, user.Banner, user.Email, user.Username, user.Password)
	defer query.Close()
}

func GetUserByEmail(email string) User {
	rows, err := DB.Query("SELECT * FROM user WHERE email='" + email + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
	}

	return user
}

func GetUserById(id int) User {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT * FROM user WHERE id = '" + id2 + "'")
	defer rows.Close()

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
	}

	return user
}

func GetUserByUuid(uuid string) User {
	rows, _ := DB.Query("SELECT * FROM user WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
	}

	return user
}

func UpdateUserInfo(user User) {
	query, err := DB.Prepare("UPDATE user set uuid = ?, profile = ?, banner = ?, username = ?, email = ?, password = ? where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(user.Uuid, user.Profile, user.Banner, user.Username, user.Email, user.Password, user.Id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was affected")
	}
}

func DeleteUser(userId int) {
	query, err := DB.Prepare("DELETE FROM user where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(userId)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was deleted")
	}
}
