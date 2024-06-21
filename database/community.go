package database

import (
	"database/sql"
	"log"
	"net/http"
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

type CommunityInfo struct {
	Id          int
	Profile     string
	Banner      string
	Name        string
	Description string
	User_id     int
	Username    string
	UserProfile string
}

func AddCommunity(community Community, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, _ := db.Prepare("INSERT INTO community (profile, banner, name, description, user_id) VALUES (?, ?, ?, ?, ?)")
	query.Exec(community.Profile, community.Banner, community.Name, community.Description, community.User_id)
	defer query.Close()
	newcommunity := GetCommunityByName(community.Name, w, r)
	AddUserCommunity(newcommunity.User_id, newcommunity.Id, w, r)
}

func GetCommunityById(id int, w http.ResponseWriter, r *http.Request) CommunityInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id2 := strconv.Itoa(id)
	rows, _ := db.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id, user.username, user.profile FROM community JOIN user ON user.id = community.user_id WHERE community.id = '" + id2 + "'")
	defer rows.Close()

	community := CommunityInfo{}

	for rows.Next() {
		rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id, &community.Username, &community.UserProfile)
	}

	return community
}

func GetCommunityBySearchString(searchString string, w http.ResponseWriter, r *http.Request) []CommunityInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id, user.username, user.profile FROM community INNER JOIN user ON user.id = community.user_id WHERE community.name LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	communityList := make([]CommunityInfo, 0)

	for rows.Next() {
		community := CommunityInfo{}
		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id, &community.Username, &community.UserProfile)
		CheckErr(err, w, r)

		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return communityList
}

func GetCommunityByName(communityName string, w http.ResponseWriter, r *http.Request) CommunityInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id, user.username, user.profile FROM community JOIN user ON user.id = community.user_id WHERE community.name = '" + communityName + "'")
	defer rows.Close()

	community := CommunityInfo{}

	for rows.Next() {
		rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id, &community.Username, &community.UserProfile)
	}

	return community
}

func GetCommunitiesByNMembers(searchString string, w http.ResponseWriter, r *http.Request) []CommunityInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	//, COUNT(user_community.community_id)
	rows, err := db.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id, user.username, user.profile FROM community JOIN user_community ON user_community.community_id = community.id JOIN user ON user.id = user_community.user_id WHERE community.name LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%' GROUP BY community.id ORDER BY COUNT(user_community.community_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	communityList := make([]CommunityInfo, 0)

	for rows.Next() {
		community := CommunityInfo{}
		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id, &community.Username, &community.UserProfile)
		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return communityList
}

func GetCommunitiesByMostPost(searchString string, w http.ResponseWriter, r *http.Request) []CommunityInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id, user.username, user.profile FROM community JOIN post ON post.community_id = community.id JOIN user ON user.id = community.user_id WHERE community.name LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%' GROUP BY community.id ORDER BY COUNT(post.community_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	communityList := make([]CommunityInfo, 0)

	for rows.Next() {
		community := CommunityInfo{}
		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id, &community.Username, &community.UserProfile)
		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return communityList
}

// func GetCommunitiesByNComment() []Community {

// 	//, COUNT(user_community.community_id)
// 	rows, err := db.Query("SELECT community.id, community.profile, community.banner, community.name, community.description, community.user_id FROM community JOIN comment ON comment.community_id = community.id GROUP BY community.id ORDER BY COUNT(comment.post_id) DESC")
// 	defer rows.Close()

// 	err = rows.Err()
// 	CheckErr(err, w, r)

// 	communityList := make([]Community, 0)

// 	for rows.Next() {
// 		community := Community{}
// 		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id)
// 		CheckErr(err, w, r)
// 		communityList = append(communityList, community)
// 	}

// 	err = rows.Err()
// 	CheckErr(err, w, r)

// 	return communityList
// }

func UpdateCommunityInfo(community Community, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE community SET profile = ?, banner = ?, name = ?, description = ?, user_id = ? WHERE id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(community.Profile, community.Banner, community.Name, community.Description, &community.User_id, community.Id)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 community was affected")
	}
}

func DeleteCommunity(communityId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM community where id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(communityId)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}

// USER_COMMUNITY Table handler

func AddUserCommunity(userId int, communityId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, _ := db.Prepare("INSERT INTO user_community (user_id, community_id) VALUES (?, ?)")
	query.Exec(userId, communityId)
	defer query.Close()
}

func ExistsUserCommunity(userId int, communityId int, w http.ResponseWriter, r *http.Request) bool {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	userid := strconv.Itoa(userId)
	communityid := strconv.Itoa(communityId)

	rows, _ := db.Query("SELECT * FROM user_community WHERE user_id = '" + userid + "' AND community_id = '" + communityid + "'")
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

func GetUsersByCommunity(communityId int, w http.ResponseWriter, r *http.Request) []User {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(communityId)
	rows, err := db.Query("SELECT user.id, user.uuid, user.profile, user.banner, user.email, user.username, user.password FROM user INNER JOIN user_community ON user.id = user_community.user_id WHERE user_community.community_id='" + id + "'")
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

func GetCommunitiesByUser(userId int, w http.ResponseWriter, r *http.Request) []Community {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(userId)
	rows, err := db.Query("SELECT * FROM community INNER JOIN user_community ON community.id = user_community.community_id WHERE user_community.user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	communityList := make([]Community, 0)

	for rows.Next() {
		community := Community{}
		err = rows.Scan(&community.Id, &community.Profile, &community.Banner, &community.Name, &community.Description, &community.User_id)
		CheckErr(err, w, r)

		communityList = append(communityList, community)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return communityList
}

func DeleteUserCommunity(userId int, communityId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM user_community WHERE user_id = ? AND community_id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(userId, communityId)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}
