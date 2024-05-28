package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Community struct {
	Id   *int
	Name string
}

func AddCommunity(community Community) {
	query, _ := DB.Prepare("INSERT INTO community (name) VALUES (?)")
	query.Exec(community.Name)
	defer query.Close()
}

func GetCommunityById(id string) Community {
	rows, _ := DB.Query("SELECT * FROM community WHERE id = '" + id + "'")
	defer rows.Close()

	community := Community{}

	for rows.Next() {
		rows.Scan(&community.Id, &community.Name)
	}

	return community
}

func DeleteCommunity(communityID string) {
	query, err := DB.Prepare("DELETE FROM community where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(communityID)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}

// USER_COMMUNITY Table handler

func AddUserCommunity(communityID string, userID string) {
	query, _ := DB.Prepare("INSERT INTO user_community (user_id, community_id) VALUES (?, ?)")
	query.Exec(communityID, userID)
	defer query.Close()
}

func GetUsersByCommunity(communityID string) []User {
	rows, err := DB.Query("SELECT * FROM user INNER JOIN user_community ON user.id = user_community.user_id WHERE user_community.community_id='" + communityID + "'")
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

func GetCommunitiesByUser(userID string) []Community {
	rows, err := DB.Query("SELECT * FROM community INNER JOIN user_community ON community.id = user_community.community_id WHERE user_community.user_id='" + userID + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	communityList := make([]Community, 0)

	for rows.Next() {
		community := Community{}
		err = rows.Scan(&community.Id, &community.Name)
		CheckErr(err)

		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err)

	return communityList
}

func DeleteUserCommunity(communityID string, userID string) {
	query, err := DB.Prepare("DELETE FROM user_community where user_id = ? AND community_id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(userID, communityID)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}
