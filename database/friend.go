package database

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

/*
!AddFriend function open data base and add friend by using the INSERT INTO sql command she take as argument two int type and a writer and request.
*/
func AddFriend(userId int, friendId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, _ := db.Prepare("INSERT INTO friend (user_id, friend_id) VALUES (?, ?)")
	query.Exec(userId, friendId)
	query.Exec(friendId, userId)
	defer query.Close()
}

/*
!ExistsFriend function open data base and check if two user are friend by using the SELECT * FROM and WHERE sql command she take as argument two int type and a writer and request and return a boolean.
*/
func ExistsFriend(userId int, friendId int, w http.ResponseWriter, r *http.Request) bool {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	userid := strconv.Itoa(userId)
	friendid := strconv.Itoa(friendId)

	rows, _ := db.Query("SELECT * FROM friend WHERE user_id = '" + userid + "' AND friend_id = '" + friendid + "'")
	defer rows.Close()

	type Friend struct {
		UserId   int
		FriendId int
	}
	friend := Friend{}

	for rows.Next() {
		rows.Scan(&friend.UserId, &friend.FriendId)
	}

	return friend != Friend{}
}

/*
!GetFriendsByUser function open data base and get friend of an user by using the SELECT * FROM and WHERE sql command she take as argument an int type and a writer and request and return a slice of user.
*/
func GetFriendsByUser(userId int, w http.ResponseWriter, r *http.Request) []User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	userid := strconv.Itoa(userId)
	rows, err := db.Query("SELECT user.id, user.uuid, user.profile, user.banner, user.email, user.username, user.password FROM user INNER JOIN friend ON user.id = friend.user_id WHERE friend.friend_id='" + userid + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	userList := make([]User, 0)

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
		CheckErr(err, w, r)

		userList = append(userList, user)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return userList
}

/*
!DeleteFriend function open data base and delete a friend relation by using the DELETE FROM and WHERE sql command she take as argument two int type and a writer and request.
*/
func DeleteFriend(userId int, friendId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM friend where user_id = ? AND friend_id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(userId, friendId)
	CheckErr(err, w, r)
	res2, err2 := query.Exec(friendId, userId)
	CheckErr(err2, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	affected2, err := res2.RowsAffected()
	CheckErr(err, w, r)

	total := strconv.FormatInt(affected+affected2, 10)
	if affected+affected2 != 2 {
		log.Fatal("Error : " + total + " elements were deleted from friend table")
	}
}
