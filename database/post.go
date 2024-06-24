package database

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Id           int
	Uuid         string
	Title        string
	Content      string
	Media        string
	MediaType    string
	User_id      int
	Community_id int
	Created      time.Time
}

type PostInfo struct {
	Id               int
	Uuid             string
	Title            string
	Content          string
	Media            string
	MediaType        string
	User_id          int
	User_uuid        string
	Username         string
	Profile          string
	Community_id     int
	CommunityName    string
	CommunityProfile string
	Created          time.Time
	Time             string
}

type TempPostInfo struct {
	Id           int
	Uuid         string
	Title        string
	Content      string
	Media        string
	MediaType    string
	User_id      int
	User_uuid    string
	Username     string
	Profile      string
	Community_id *int
	Created      time.Time
	Time         string
}

/*
!AddUser function open data base and add an psot to it with the INSERT INTO sql command she take as argument an Post type and a writer and request.
*/
func AddPost(post Post, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("INSERT INTO post (uuid, title, content, media, media_type, user_id, community_id, created) VALUES (?, ?, ?, ?, ?, ?, NULLIF(?, 0), ?)")
	CheckErr(err, w, r)
	query.Exec(post.Uuid, post.Title, post.Content, post.Media, post.MediaType, post.User_id, post.Community_id, post.Created)
	defer query.Close()
}

/*
!GetPostsBySearchString function open data base and get post by search string by using the SELECT sql command she take as argument a string type, a writer, a request and return a sice of PostInfo.
*/
func GetPostsBySearchString(searchString string, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.title LIKE '%" + searchString + "%' OR post.content LIKE '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return postList
}

/*
!GetPostsByUser function open data base and get post by an user id by using the SELECT sql command she take as argument an int type, a writer, a request and return a sice of PostInfo.
*/
func GetPostsByUser(userId int, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(userId)
	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return postList
}

/*
!GetPostsByCommunity function open data base and get post by a community id by using the SELECT sql command she take as argument an int type, a writer, a request and return a sice of PostInfo.
*/
func GetPostsByCommunity(communityId int, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(communityId)
	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.community_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return postList
}

/*
!GetPostById function open data base and get post by an id by using the SELECT sql command she take as argument an int type, a writer, a request and return a PostInfo.
*/
func GetPostById(id int, w http.ResponseWriter, r *http.Request) PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id2 := strconv.Itoa(id)
	rows, _ := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.id = '" + id2 + "'")
	defer rows.Close()

	temppostInfo := TempPostInfo{}

	for rows.Next() {
		rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
	}

	postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
	if temppostInfo.Community_id != nil {
		postInfo.Community_id = *temppostInfo.Community_id
		community := GetCommunityById(*temppostInfo.Community_id, w, r)
		postInfo.CommunityName = community.Name
		postInfo.CommunityProfile = community.Profile
	}

	return postInfo
}

/*
!GetPostByUuid function open data base and get post by an uuid by using the SELECT sql command she take as argument a string type, a writer, a request and return a PostInfo.
*/
func GetPostByUuid(uuid string, w http.ResponseWriter, r *http.Request) PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, _ := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.uuid = '" + uuid + "'")
	defer rows.Close()

	temppostInfo := TempPostInfo{}

	for rows.Next() {
		rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
	}

	postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
	if temppostInfo.Community_id != nil {
		postInfo.Community_id = *temppostInfo.Community_id
		community := GetCommunityById(*temppostInfo.Community_id, w, r)
		postInfo.CommunityName = community.Name
		postInfo.CommunityProfile = community.Profile
	}

	return postInfo
}

/*
!GetAllPosts function open data base and get all existing posts using the SELECT sql command she take as argument a writer, a request and return a slice of PostInfo.
*/
func GetAllPosts(w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id JOIN community ON community.id = post.community_id JOIN like ON like.post_id = post.id")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return postList
}

/*
!GetPostByMostComment function open data base and sort post by most commented and by searched string by using the SELECT sql command she take as argument a string type, a writer, a request and return a slice of PostInfo.
*/
func GetPostByMostComment(searchString string, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id JOIN comment ON comment.post_id = post.id WHERE post.title LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%' GROUP BY post.id ORDER BY COUNT(comment.post_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	//Get the posts with 0 Commments that were not sent by the pervious sql query
	allrows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id WHERE post.title LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%'")
	defer allrows.Close()

	err = allrows.Err()
	CheckErr(err, w, r)

	allPostList := make([]PostInfo, 0)

	for allrows.Next() {
		temppostInfo := TempPostInfo{}
		err = allrows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		allPostList = append(allPostList, postInfo)
	}

	err = allrows.Err()
	CheckErr(err, w, r)

	for i := 0; i < len(allPostList); i++ {
		if !ContainsPost(postList, allPostList[i]) {
			postList = append(postList, allPostList[i])
		}
	}

	return postList
}

/*
!GetPostByPopular function open data base and sort post by most followed and by searched string by using the SELECT sql command she take as argument a string type, a writer, a request and return a slice of PostInfo.
*/
func GetPostByPopular(searchString string, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id JOIN like ON like.post_id = post.id WHERE like.rating = 'like' AND (post.title LIKE '%" + searchString + "%' OR post.content LIKE '%" + searchString + "%') GROUP BY post.id ORDER BY COUNT(like.post_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return postList
}

/*
!GetPostByPopularByCommunity function open data base and sort post by most followed in a given community by using the SELECT sql command she take as argument an int type, a writer, a request and return a slice of PostInfo.
*/
func GetPostByPopularByCommunity(communityId int, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	communityid := strconv.Itoa(communityId)
	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id JOIN community ON community.id = post.community_id JOIN like ON like.post_id = post.id WHERE like.rating = 'like' AND post.community_id = " + communityid + " GROUP BY post.id ORDER BY COUNT(like.post_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return postList
}

/*
!GetPostByMostCommentByCommunity function open data base and sort post by most commented in a given community by using the SELECT sql command she take as argument an int type, a writer, a request and return a slice of PostInfo.
*/
func GetPostByMostCommentByCommunity(communityId int, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	communityid := strconv.Itoa(communityId)
	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id JOIN community ON community.id = post.community_id JOIN comment ON comment.post_id = post.id WHERE post.community_id = " + communityid + " GROUP BY post.id ORDER BY COUNT(comment.post_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	//Get the posts with 0 Commments that were not sent by the pervious sql query
	allrows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id JOIN community ON community.id = post.community_id WHERE post.community_id = " + communityid)
	defer allrows.Close()

	err = allrows.Err()
	CheckErr(err, w, r)

	allPostList := make([]PostInfo, 0)

	for allrows.Next() {
		temppostInfo := TempPostInfo{}
		err = allrows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		allPostList = append(allPostList, postInfo)
	}

	err = allrows.Err()
	CheckErr(err, w, r)

	for i := 0; i < len(allPostList); i++ {
		if !ContainsPost(postList, allPostList[i]) {
			postList = append(postList, allPostList[i])
		}
	}

	return postList
}

func GetPostByPopularByUser(userId int, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	userid := strconv.Itoa(userId)
	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id JOIN like ON like.post_id = post.id WHERE like.rating = 'like' AND post.user_id = " + userid + " GROUP BY post.id ORDER BY COUNT(like.post_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return postList
}

func GetPostByMostCommentByUser(userId int, w http.ResponseWriter, r *http.Request) []PostInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	userid := strconv.Itoa(userId)
	rows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id JOIN comment ON comment.post_id = post.id WHERE post.user_id = " + userid + " GROUP BY post.id ORDER BY COUNT(comment.post_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	//Get the posts with 0 Commments that were not sent by the pervious sql query
	allrows, err := db.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.uuid, user.username, user.profile, post.community_id, post.created FROM post JOIN user ON user.id = post.user_id WHERE post.user_id = " + userid)
	defer allrows.Close()

	err = allrows.Err()
	CheckErr(err, w, r)

	allPostList := make([]PostInfo, 0)

	for allrows.Next() {
		temppostInfo := TempPostInfo{}
		err = allrows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.User_uuid, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err, w, r)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, User_uuid: temppostInfo.User_uuid, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			community := GetCommunityById(*temppostInfo.Community_id, w, r)
			postInfo.CommunityName = community.Name
			postInfo.CommunityProfile = community.Profile
		}
		allPostList = append(allPostList, postInfo)
	}

	err = allrows.Err()
	CheckErr(err, w, r)

	for i := 0; i < len(allPostList); i++ {
		if !ContainsPost(postList, allPostList[i]) {
			postList = append(postList, allPostList[i])
		}
	}

	return postList
}

/*
!UpdatePostInfo function open data base and update post information by using the SELECT sql command she take as argument an Post type, a writer, a request.
*/
func UpdatePostInfo(post Post, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE post SET uuid = ?, title = ?, content = ?, media = ?, media_type = ?, user_id = ?, community_id = NULLIF(?, 0), created = ? WHERE id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(post.Uuid, post.Title, post.Content, post.Media, post.MediaType, post.User_id, post.Community_id, post.Created, post.Id)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was affected")
	}
}

/*
!DeletePost function open data base and delete post information by using the DELETE sql command she take as argument an int type, a writer, a request.
*/
func DeletePost(postId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM post where id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(postId)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was deleted")
	}
}
