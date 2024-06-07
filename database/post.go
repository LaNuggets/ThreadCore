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
	User_id      int
	Community_id int
	Created      time.Time
}

type PostInfo struct {
	Id           int
	Uuid         string
	Title        string
	Content      string
	Media        string
	User_id      int
	Username     string
	Community_id int
	CommuntyName string
	Created      time.Time
}

type TempPostInfo struct {
	Id           int
	Uuid         string
	Title        string
	Content      string
	Media        string
	User_id      int
	Username     string
	Community_id *int
	Created      time.Time
}

func AddPost(post Post) {
	query, _ := DB.Prepare("INSERT INTO post (uuid,title, content, media, user_id, community_id, created) VALUES (?, ?, ?, ?, ?, NULLIF(?, 0), ?)")
	query.Exec(post.Uuid, post.Title, post.Content, post.Media, post.User_id, post.Community_id, post.Created)
	defer query.Close()
}

func GetPostsBySearchString(searchString string) []PostInfo {
	rows, err := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.user_id, user.username, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.title LIKE '%" + searchString + "%' OR post.content LIKE '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Community_id: 0, CommuntyName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			postInfo.CommuntyName = GetCommunityById(*temppostInfo.Community_id).Name
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostsByUser(userId int) []PostInfo {
	id := strconv.Itoa(userId)
	rows, err := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.user_id, user.username, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Community_id: 0, CommuntyName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			postInfo.CommuntyName = GetCommunityById(*temppostInfo.Community_id).Name
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostsByCommunity(communityId int) []PostInfo {
	id := strconv.Itoa(communityId)
	rows, err := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.user_id, user.username, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.community_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostInfo, 0)

	for rows.Next() {
		temppostInfo := TempPostInfo{}
		err = rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Community_id, &temppostInfo.Created)
		CheckErr(err)
		postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Community_id: 0, CommuntyName: "", Created: temppostInfo.Created}
		if temppostInfo.Community_id != nil {
			postInfo.Community_id = *temppostInfo.Community_id
			postInfo.CommuntyName = GetCommunityById(*temppostInfo.Community_id).Name
		}
		postList = append(postList, postInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostById(id int) PostInfo {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.user_id, user.username, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.id = '" + id2 + "'")
	defer rows.Close()

	temppostInfo := TempPostInfo{}

	for rows.Next() {
		rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Community_id, &temppostInfo.Created)

	}

	postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Community_id: 0, CommuntyName: "", Created: temppostInfo.Created}
	if temppostInfo.Community_id != nil {
		postInfo.Community_id = *temppostInfo.Community_id
		postInfo.CommuntyName = GetCommunityById(*temppostInfo.Community_id).Name
	}

	return postInfo
}

func GetPostByUuid(uuid string) PostInfo {
	rows, _ := DB.Query("SELECT post.id, post.uuid, post.title, post.content, post.media, post.user_id, user.username, post.community_id, post.created FROM post INNER JOIN user ON user.id = post.user_id WHERE post.uuid = '" + uuid + "'")
	defer rows.Close()

	temppostInfo := TempPostInfo{}

	for rows.Next() {
		rows.Scan(&temppostInfo.Id, &temppostInfo.Uuid, &temppostInfo.Title, &temppostInfo.Content, &temppostInfo.Media, &temppostInfo.User_id, &temppostInfo.Username, &temppostInfo.Community_id, &temppostInfo.Created)
	}

	postInfo := PostInfo{Id: temppostInfo.Id, Uuid: temppostInfo.Uuid, Title: temppostInfo.Title, Content: temppostInfo.Content, Media: temppostInfo.Media, User_id: temppostInfo.User_id, Username: temppostInfo.Username, Community_id: 0, CommuntyName: "", Created: temppostInfo.Created}
	if temppostInfo.Community_id != nil {
		postInfo.Community_id = *temppostInfo.Community_id
		postInfo.CommuntyName = GetCommunityById(*temppostInfo.Community_id).Name
	}

	return postInfo
}

func UpdatePostInfo(post Post) {
	query, err := DB.Prepare("UPDATE post set title = ?, content = ?, media = ?, user_id = ?, community_id = ?, created = ? where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(post.Title, post.Content, post.Media, post.User_id, post.Community_id, post.Created, post.Id)
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
