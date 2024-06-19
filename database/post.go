package database

import (
	"log"
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
	Id            int
	Uuid          string
	Title         string
	Content       string
	Media         string
	MediaType     string
	User_id       int
	Username      string
	Profile       string
	Community_id  int
	CommunityName string
	Created       time.Time
	Time          string
}

type TempPostInfo struct {
	Id           int
	Uuid         string
	Title        string
	Content      string
	Media        string
	MediaType    string
	User_id      int
	Username     string
	Profile      string
	Community_id *int
	Created      time.Time
	Time         string
}

func AddPost(post Post) {
	query, err := DB.Prepare("INSERT INTO post (uuid, title, content, media, media_type, user_id, community_id, created) VALUES (?, ?, ?, ?, ?, ?, NULLIF(?, 0), ?)")
	CheckErr(err)
	query.Exec(post.Uuid, post.Title, post.Content, post.Media, post.MediaType, post.User_id, post.Community_id, post.Created)
	defer query.Close()
}

func GetPostsBySearchString(searchString string) []PostInfo {
	rows, err := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.title LIKE '%" + searchString + "%' OR post.content LIKE '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			postInfo.CommunityName = GetCommunityById(*temppostInfo.Community_id).Name
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostsByUser(userId int) []PostInfo {
	id := strconv.Itoa(userId)
	rows, err := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			postInfo.CommunityName = GetCommunityById(*temppostInfo.Community_id).Name
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostsByCommunity(communityId int) []PostInfo {
	id := strconv.Itoa(communityId)
	rows, err := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.community_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			postInfo.CommunityName = GetCommunityById(*temppostInfo.Community_id).Name
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostById(id int) PostInfo {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type post.user_id, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.id = '" + id2 + "'")
	defer rows.Close()

	temppostInfo := TempPostInfo{}

	for rows.Next() {
		rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)

	}

	postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
	if temppostInfo.Community_id != nil {
		postInfo.Community_id = *temppostInfo.Community_id
		postInfo.CommunityName = GetCommunityById(*temppostInfo.Community_id).Name
	}

	return postInfo
}

func GetPostByUuid(uuid string) PostInfo {
	rows, _ := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.media_type, post.user_id, user.username, user.profile, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.uuid = '" + uuid + "'")
	defer rows.Close()

	temppostInfo := TempPostInfo{}

	for rows.Next() {
		rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
	}

	postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
	if temppostInfo.Community_id != nil {
		postInfo.Community_id = *temppostInfo.Community_id
		postInfo.CommunityName = GetCommunityById(*temppostInfo.Community_id).Name
	}

	return postInfo
}

func GetPostByMostComment(searchString string) []PostInfo {

	rows, err := DB.Query("SELECT post.id, post.title, post.content, post.media, post.media_type, post.user_id, user.username, user.profile, post.community_id, community.name, post.created FROM post JOIN user ON user.id = post.user_id JOIN community ON community.id = post.community_id JOIN comment ON comment.post_id = post.id WHERE post.title LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%' GROUP BY post.id ORDER BY COUNT(comment.post_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			postInfo.CommunityName = GetCommunityById(*temppostInfo.Community_id).Name
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostByPopular(searchString string) []PostInfo {

	rows, err := DB.Query("SELECT post.id, post.title, post.content, post.media, post.media_type, post.user_id, user.username, post.community_id, community.name, post.created FROM post JOIN user ON user.id = post.user_id JOIN community ON community.id = post.community_id JOIN like ON like.post_id = post.id WHERE like.rating = 'like' AND post.title LIKE '%" + searchString + "%' OR user.username LIKE '%" + searchString + "%' GROUP BY post.id ORDER BY COUNT(like.post_id) DESC")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.MediaType, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Profile, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, MediaType: temppostInfo.MediaType, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Profile: temppostInfo.Profile, Community_id: 0, CommunityName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			postInfo.CommunityName = GetCommunityById(*temppostInfo.Community_id).Name
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func UpdatePostInfo(post Post) {
	query, err := DB.Prepare("UPDATE post set title = ?, content = ?, media = ?, media_type = ?, user_id = ?, community_id = ?, created = ? where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(post.Title, post.Content, post.Media, post.MediaType, post.User_id, post.Community_id, post.Created, post.Id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was affected")
	}
}

func DeletePost(postId int) {
	query, err := DB.Prepare("DELETE FROM post where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(postId)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was deleted")
	}
}
