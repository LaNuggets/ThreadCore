package database

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Community struct {
	Id          int
	Profile     string
	Banner      string
	Name        string
	Description string
	User_id     int
}

type CommunityDisplay struct {
	Id          int
	Profile     string
	Banner      string
	Name        string
	Description string
	User_id     int
	Username    string
}

func AddCommunity(community Community) {
	query, _ := DB.Prepare("INSERT INTO community (profile, banner, name, description, user_id) VALUES (?, ?, ?, ?, ?)")
	query.Exec(community.Profile, community.Banner, community.Name, community.Description, community.User_id)
	defer query.Close()
	newcommunity := GetCommunityByName(community.Name)
	AddUserCommunity(newcommunity.User_id, newcommunity.Id)
}

func GetCommunityById(id int) Community {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT * FROM community WHERE id = '" + id2 + "'")
	defer rows.Close()

	community := Community{}

	for rows.Next() {
		rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id)
	}

	return community
}

func GetCommunityBySearchString(searchString string) []CommunityDisplay {
	rows, err := DB.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id, user.username FROM community INNER JOIN user ON user.id = community.user_id WHERE community.name LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	communityList := make([]CommunityDisplay, 0)

	for rows.Next() {
		community := CommunityDisplay{}
		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id, &community.Username)
		CheckErr(err)

		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err)

	return communityList
}

func GetCommunityByName(communityName string) Community {
	rows, _ := DB.Query("SELECT * FROM community WHERE name = '" + communityName + "'")
	defer rows.Close()

	community := Community{}

	for rows.Next() {
		rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id)
	}

	return community
}

func GetCommunitiesByNMembers(searchString string) []CommunityDisplay {

	//, COUNT(user_community.community_id)
	rows, err := DB.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id, user.username FROM community JOIN user_community ON user_community.community_id = community.id JOIN user ON user.id = user_community.user_id WHERE community.name LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%' GROUP BY community.id ORDER BY COUNT(user_community.community_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	communityList := make([]CommunityDisplay, 0)

	for rows.Next() {
		community := CommunityDisplay{}
		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id, &community.Username)
		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err)

	return communityList
}

func GetCommunitiesByMostPost(searchString string) []CommunityDisplay {

	rows, err := DB.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id, user.username FROM community JOIN post ON post.community_id = community.id JOIN user ON user.id = community.user_id WHERE community.name LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%' GROUP BY community.id ORDER BY COUNT(post.community_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	communityList := make([]CommunityDisplay, 0)

	for rows.Next() {
		community := CommunityDisplay{}
		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id, &community.Username)
		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err)

	return communityList
}

// func GetCommunitiesByNComment() []Community {

// 	//, COUNT(user_community.community_id)
// 	rows, err := DB.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id FROM community JOIN comment ON comment.community_id = community.id GROUP BY community.id ORDER BY COUNT(comment.post_id) DESC")
// 	defer rows.Close()

// 	err = rows.Err()
// 	CheckErr(err)

// 	communityList := make([]Community, 0)

// 	for rows.Next() {
// 		community := Community{}
// 		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id)
// 		CheckErr(err)
// 		communityList = append(communityList, community)
// 	}

// 	err = rows.Err()
// 	CheckErr(err)

// 	return communityList
// }

func UpdateCommunityInfo(community Community) {
	query, err := DB.Prepare("UPDATE community SET profile = ?, banner = ?, name = ?, description = ?, user_id = ? WHERE id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(community.Profile, community.Banner, community.Name, community.Description, &community.User_id, community.Id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 community was affected")
	}
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
	query.Exec(userId, communityId)
	defer query.Close()
}

func ExistsUserCommunity(userId int, communityId int) bool {
	userid := strconv.Itoa(userId)
	communityid := strconv.Itoa(communityId)

	rows, _ := DB.Query("SELECT * FROM user_community WHERE user_id = '" + userid + "' AND community_id = '" + communityid + "'")
	defer rows.Close()

	type UserCommunity struct {
		UserId      int
		CommunityId int
	}
	user_communty := UserCommunity{}

	for rows.Next() {
		rows.Scan(&user_communty.UserId, &user_communty.CommunityId)
	}

	return user_communty != UserCommunity{}
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
		err = rows.Scan(&user.Id, &user.Uuid, &user.Profile, &user.Banner, &user.Email, &user.Username, &user.Password)
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
		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id)
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

	res, err := query.Exec(userId, communityId)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}
