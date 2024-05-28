package database

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func AddFriend(userID string, friendID string) {
	query, _ := DB.Prepare("INSERT INTO friend (user_id, friend_id) VALUES (?, ?)")
	query.Exec(userID, friendID)
	query.Exec(friendID, userID)
	defer query.Close()
}

func GetFriendsByUser(userId string) []User {
	rows, err := DB.Query("SELECT * FROM user INNER JOIN friend ON user.id = friend.user_id WHERE friend.friend_id='" + userId + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	userList := make([]User, 0)

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.ProfilePicture, &user.Email, &user.Username, &user.Password)
		CheckErr(err)

		userList = append(userList, user)
	}

	err = rows.Err()
	CheckErr(err)

	return userList
}

func DeleteFriend(userID string, friendID string) {
	query, err := DB.Prepare("DELETE FROM friend where user_id = ? AND friend_id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(userID, friendID)
	CheckErr(err)
	res2, err2 := query.Exec(friendID, userID)
	CheckErr(err2)

	affected, err := res.RowsAffected()
	CheckErr(err)

	affected2, err := res2.RowsAffected()
	CheckErr(err)

	total := strconv.FormatInt(affected+affected2, 10)
	if affected+affected2 != 2 {
		log.Fatal("Error : " + total + " elements were deleted from friend table")
	}
}