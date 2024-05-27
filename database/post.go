package database

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	id           string
	title        string
	content      string
	user_id      string
	community_id string
	created      time.Time
}

func AddPost(post Post) {
	id := ""      // ADD UUID
	created := "" //ADD time
	query, _ := DB.Prepare("INSERT INTO post (id, title, content, user_id, community_id, created) VALUES (?, ?, ?, ?, ?)")
	query.Exec(id, post.title, post.content, post.user_id, post.community_id, created)
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
		err = rows.Scan(&post.id, &post.title, &post.content, &post.user_id, &post.community_id, &post.created)
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
		err = rows.Scan(&post.id, &post.title, &post.content, &post.user_id, &post.community_id, &post.created)
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
		err = rows.Scan(&post.id, &post.title, &post.content, &post.user_id, &post.community_id, &post.created)
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
		rows.Scan(&post.id, &post.title, &post.content, &post.user_id, &post.community_id, &post.created)
	}

	return post
}

func UpdatePostInfo(post Post) {
	stmt, err := DB.Prepare("UPDATE post set title = ?, content = ?, user_id = ?, community_id = ?, created = ? where id = ?")
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(post.title, post.content, post.user_id, post.community_id, post.created, post.id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was affected")
	}
}

func DeletePost(userID string) {
	stmt, err := DB.Prepare("DELETE FROM user where id = ?")
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(userID)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was deleted")
	}
}
