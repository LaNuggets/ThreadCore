package database

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Community struct {
	Id        int
	Name      string
	Following int
}

func AddCommunity(community Community) {
	query, _ := DB.Prepare("INSERT INTO community (name, following) VALUES (?, ?)")
	query.Exec(community.Name, 0)
	defer query.Close()
}

func GetCommunityById(id int) Community {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT * FROM community WHERE id = '" + id2 + "'")
	defer rows.Close()

	community := Community{}

	for rows.Next() {
		rows.Scan(&community.Id, &community.Name, &community.Following)
	}

	return community
}

func GetCommunityByName(communityName string) Community {
	rows, _ := DB.Query("SELECT * FROM community WHERE name = '" + communityName + "'")
	defer rows.Close()

	community := Community{}

	for rows.Next() {
		rows.Scan(&community.Id, &community.Name, &community.Following)
	}

	return community
}

func GetCommunitiesByNMembers() []Community {
	rows, err := DB.Query("SELECT * FROM community ORDER BY following DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	communityList := make([]Community, 0)

	for rows.Next() {
		community := Community{}
		err = rows.Scan(&community.Id, &community.Name, &community.Following)
		CheckErr(err)

		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err)

	return communityList
}

func DeleteCommunity(communityId int) {
	query, err := DB.Prepare("DELETE FROM community where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(communityId)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}

// USER_COMMUNITY Table handler

func AddUserCommunity(userId int, communityId int) {
	query, _ := DB.Prepare("INSERT INTO user_community (user_id, community_id) VALUES (?, ?)")
	query.Exec(communityId, userId)
	defer query.Close()

	query2, _ := DB.Prepare("UPDATE community SET following=following + 1 WHERE id = ?")
	query2.Exec(communityId)
	defer query2.Close()
}

func GetUsersByCommunity(communityId int) []User {
	id := strconv.Itoa(communityId)
	rows, err := DB.Query("SELECT * FROM user INNER JOIN user_community ON user.id = user_community.user_id WHERE user_community.community_id='" + id + "'")
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

func GetCommunitiesByUser(userId int) []Community {
	id := strconv.Itoa(userId)
	rows, err := DB.Query("SELECT * FROM community INNER JOIN user_community ON community.id = user_community.community_id WHERE user_community.user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	communityList := make([]Community, 0)

	for rows.Next() {
		community := Community{}
		err = rows.Scan(&community.Id, &community.Name, &community.Following)
		CheckErr(err)

		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err)

	return communityList
}

func DeleteUserCommunity(communityId int, userId int) {
	query, err := DB.Prepare("DELETE FROM user_community where user_id = ? AND community_id = ?")
	CheckErr(err)
	defer query.Close()

	query2, _ := DB.Prepare("UPDATE community SET following = following - 1 WHERE id = ?")
	query2.Exec(communityId)
	defer query2.Close()

	res, err := query.Exec(userId, communityId)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}
