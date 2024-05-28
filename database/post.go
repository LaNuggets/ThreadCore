package database

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Id           *int
	Title        string
	Content      string
	User_id      int
	Community_id int
	Created      time.Time
}

func AddPost(post Post) {
	query, _ := DB.Prepare("INSERT INTO post (title, content, user_id, community_id, created) VALUES (?, ?, ?, ?, ?)")
	query.Exec(post.Title, post.Content, post.User_id, post.Community_id, post.Created)
	defer query.Close()
}

func GetPostsBySearchString(searchString string) []Post {
	rows, err := DB.Query("SELECT * FROM post WHERE title like '%" + searchString + "%' OR content like '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]Post, 0)

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.User_id, &post.Community_id, &post.Created)
		CheckErr(err)

		postList = append(postList, post)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostsByUser(userID string) []Post {
	rows, err := DB.Query("SELECT * FROM post WHERE user_id='" + userID + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]Post, 0)

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.User_id, &post.Community_id, &post.Created)
		CheckErr(err)

		postList = append(postList, post)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostsByCommunity(communityID string) []Post {
	rows, err := DB.Query("SELECT * FROM post WHERE community_id='" + communityID + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	postList := make([]Post, 0)

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.User_id, &post.Community_id, &post.Created)
		CheckErr(err)

		postList = append(postList, post)
	}

	err = rows.Err()
	CheckErr(err)

	return postList
}

func GetPostById(id string) Post {
	rows, _ := DB.Query("SELECT * FROM post WHERE id = '" + id + "'")
	defer rows.Close()

	post := Post{}

	for rows.Next() {
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.User_id, &post.Community_id, &post.Created)
	}

	return post
}

func UpdatePostInfo(post Post) {
	query, err := DB.Prepare("UPDATE post set title = ?, content = ?, user_id = ?, community_id = ?, created = ? where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(post.Title, post.Content, post.User_id, post.Community_id, post.Created, post.Id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was affected")
	}
}

func DeletePost(postID string) {
	query, err := DB.Prepare("DELETE FROM post where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(postID)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was deleted")
	}
}
