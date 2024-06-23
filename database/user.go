package database

import (
	"database/sql"
	"log"
	"net/http"
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

/*
!AddUser function open data base and add an user to it with the INSERT INTO sql command she take as argument an User type and a writer and request.
*/
func AddUser(user User, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, _ := db.Prepare("INSERT INTO user (uuid, profile, banner, email, username, password) VALUES (?, ?, ?, ?, ?, ?)")
	query.Exec(user.Uuid, user.Profile, user.Banner, user.Email, user.Username, user.Password)
	defer query.Close()

	query2, _ := db.Prepare("INSERT INTO friend (user_id, friend_id) VALUES (?, ?)")
	query2.Exec(user.Id, user.Id)
	defer query2.Close()
}

/*
!GetUserByEmail function is used to get a user by is email by using the SELECT * FROM sql command. She take as argument a string, a writer,a request and return an User type.
*/
func GetUserByEmail(email string, w http.ResponseWriter, r *http.Request) User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE email='" + email + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
	}

	return user
}

/*
!GetUserById function is used to get a user by is id by using the SELECT * FROM sql command. She take as argument a int, a writer, a request and return an User type.
*/
func GetUserById(id int, w http.ResponseWriter, r *http.Request) User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id2 := strconv.Itoa(id)
	rows, _ := db.Query("SELECT * FROM user WHERE id = '" + id2 + "'")
	defer rows.Close()

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
	}

	return user
}

/*
!GetUserByUuid function is used to get a user by is uuid by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return an User type.
*/
func GetUserByUuid(uuid string, w http.ResponseWriter, r *http.Request) User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM user WHERE uuid = '" + uuid + "'")
	defer rows.Close()

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
	}

	return user
}

/*
!GetUserByUsername function is used to get a user by is username by using the SELECT * FROM sql command. She take as argument a string, a writer, a request and return an User type.
*/
func GetUserByUsername(username string, w http.ResponseWriter, r *http.Request) User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE username='" + username + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	user := User{}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
	}

	return user
}

/*
!GetUserBySearchString function is used to get a user by a searched string by using the SELECT * FROM and WHERE LIKE sql command. She take as argument a string, a writer, a request and return an slice of User type.
*/
func GetUserBySearchString(searchString string, w http.ResponseWriter, r *http.Request) []User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE username LIKE '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	userList := make([]User, 0)

	for rows.Next() {
		userDisplay := User{}
		err = rows.Scan(&userDisplay.Id, &userDisplay.Uuid, &userDisplay.Profile, &userDisplay.Banner, &userDisplay.Email, &userDisplay.Username, &userDisplay.Password)
		CheckErr(err, w, r)

		userList = append(userList, userDisplay)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return userList
}

/*
!GetUserByMostPopular function is used to sort user by most popular and by a searched string by using SELECT * FROM and WHERE LIKE sql command. She take as argument a string, a writer, a request and return an slice of User type.
*/
func GetUserByMostPopular(searchString string, w http.ResponseWriter, r *http.Request) []User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT user.id, user.uuid, user.profile, user.banner, user.email, user.username, user.password FROM user JOIN friend ON friend.user_id = user.id WHERE username LIKE '%" + searchString + "%' GROUP BY user.id ORDER BY COUNT(friend.user_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	userList := make([]User, 0)

	for rows.Next() {
		userDisplay := User{}
		err = rows.Scan(&userDisplay.Id, &userDisplay.Uuid, &userDisplay.Profile, &userDisplay.Banner, &userDisplay.Email, &userDisplay.Username, &userDisplay.Password)
		CheckErr(err, w, r)

		userList = append(userList, userDisplay)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return userList
}

/*
!GetUserByMostPost function is used to sort user by most post they did and by a searched string by using SELECT * FROM and WHERE LIKE sql command. She take as argument a string, a writer, a request and return an slice of User type.
*/
func GetUserByMostPost(searchString string, w http.ResponseWriter, r *http.Request) []User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT user.id, user.uuid, user.profile, user.banner, user.email, user.username, user.password FROM user JOIN post ON post.user_id = user.id WHERE username LIKE '%" + searchString + "%' GROUP BY user.id ORDER BY COUNT(post.user_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	userList := make([]User, 0)

	for rows.Next() {
		userDisplay := User{}
		err = rows.Scan(&userDisplay.Id, &userDisplay.Uuid, &userDisplay.Profile, &userDisplay.Banner, &userDisplay.Email, &userDisplay.Username, &userDisplay.Password)
		CheckErr(err, w, r)

		userList = append(userList, userDisplay)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return userList
}

/*
!GetUserByMostComment function is used to sort user by most comment they did and by a searched string by using SELECT * FROM and WHERE LIKE sql command. She take as argument a string, a writer, a request and return an slice of User type.
*/
func GetUserByMostComment(searchString string, w http.ResponseWriter, r *http.Request) []User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT user.id, user.uuid, user.profile, user.banner, user.email, user.username, user.password FROM user JOIN comment ON comment.user_id = user.id WHERE username LIKE '%" + searchString + "%' GROUP BY user.id ORDER BY COUNT(comment.user_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	userList := make([]User, 0)

	for rows.Next() {
		userDisplay := User{}
		err = rows.Scan(&userDisplay.Id, &userDisplay.Uuid, &userDisplay.Profile, &userDisplay.Banner, &userDisplay.Email, &userDisplay.Username, &userDisplay.Password)
		CheckErr(err, w, r)

		userList = append(userList, userDisplay)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return userList
}

/*
!UpdateUserInfo function is used to update user inforamation by using UPDATE sql command. She take as argument a user type, a writer, a request.
*/
func UpdateUserInfo(user User, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE user SET uuid = ?, profile = ?, banner = ?, username = ?, email = ?, password = ? WHERE id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(user.Uuid, user.Profile, user.Banner, user.Username, user.Email, user.Password, user.Id)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was affected")
	}
}

/*
!DeleteUser function is used to delete user by using DELETE sql command. She take as argument an int, a writer, a request.
*/
func DeleteUser(userId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM user WHERE id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(userId)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was deleted")
	}
}
