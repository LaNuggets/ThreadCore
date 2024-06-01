package database

import (
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Id           int
	Title        string
	Content      string
	Media        string
	User_id      int
	Community_id int
	Created      time.Time
}

type PostDisplay struct {
	Id           int
	Title        string
	Content      string
	Media        string
	User_id      int
	Community_id int
	Created      time.Time
	Username     string
}

func AddPost(post Post) {
	query, _ := DB.Prepare("INSERT INTO post (title, content, media, user_id, community_id, created) VALUES (?, ?, ?, ?, NULLIF(?, 0), ?)")
	query.Exec(post.Title, post.Content, post.Media, post.User_id, post.Community_id, post.Created)
	defer query.Close()
}

func GetPostsBySearchString(searchString string) []PostDisplay {
	rows, err := DB.Query("SELECT post.id, post.title, post.content, post.media, post.user_id, post.community_id, post.created, user.username FROM post INNER JOIN user ON user.id = post.user_id WHERE title LIKE '%" + searchString + "%' OR content LIKE '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostDisplay, 0)

	for rows.Next() {
		postDisplay := PostDisplay{}
		err = rows.Scan(&postDisplay.Id, &postDisplay.Title, &postDisplay.Content, &postDisplay.Media, &postDisplay.User_id, &postDisplay.Community_id, &postDisplay.Created, &postDisplay.Username)
		CheckErr(err)

		postList = append(postList, postDisplay)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostsByUser(userId int) []Post {
	id := strconv.Itoa(userId)
	rows, err := DB.Query("SELECT * FROM post WHERE user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]Post, 0)

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.Media, &post.User_id, &post.Community_id, &post.Created)
		CheckErr(err)

		postList = append(postList, post)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostsByCommunity(communityId int) []PostDisplay {
	id := strconv.Itoa(communityId)
	rows, err := DB.Query("SELECT post.id, post.title, post.content, post.media, post.user_id, post.community_id, post.created, user.username FROM post INNER JOIN user ON user.id = post.user_id WHERE community_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]PostDisplay, 0)

	for rows.Next() {
		postDisplay := PostDisplay{}
		err = rows.Scan(&postDisplay.Id, &postDisplay.Title, &postDisplay.Content, &postDisplay.Media, &postDisplay.User_id, &postDisplay.Community_id, &postDisplay.Created, &postDisplay.Username)
		CheckErr(err)

		postList = append(postList, postDisplay)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostById(id int) Post {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT * FROM post WHERE id = '" + id2 + "'")
	defer rows.Close()

	post := Post{}

	for rows.Next() {
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.Media, &post.User_id, &post.Community_id, &post.Created)
	}

	return post
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
