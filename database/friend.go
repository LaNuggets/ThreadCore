package database

import (
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func AddFriend(userId int, friendId int) {
	query, _ := DB.Prepare("INSERT INTO friend (user_id, friend_id) VALUES (?, ?)")
	query.Exec(userId, friendId)
	query.Exec(friendId, userId)
	defer query.Close()
}

func ExistsFriend(userId int, friendId int, w http.ResponseWriter, r *http.Request) bool {
	userid := strconv.Itoa(userId)
	friendid := strconv.Itoa(friendId)

	rows, _ := DB.Query("SELECT * FROM friend WHERE user_id = '" + userid + "' AND friend_id = '" + friendid + "'")
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

func GetFriendsByUser(userId int, w http.ResponseWriter, r *http.Request) []User {
	userid := strconv.Itoa(userId)
	rows, err := DB.Query("SELECT user.id, user.uuid, user.profile, user.banner, user.email, user.username, user.password FROM user INNER JOIN friend ON user.id = friend.user_id WHERE friend.friend_id='" + userid + "'")
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

func DeleteFriend(userId int, friendId int, w http.ResponseWriter, r *http.Request) {
	query, err := DB.Prepare("DELETE FROM friend where user_id = ? AND friend_id = ?")
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
