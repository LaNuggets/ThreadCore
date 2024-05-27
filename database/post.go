package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type post struct {
	id           string
	title        string
	content      string
	user_id      string
	community_id string
	created      time.Time
}

func addPost(db *sql.DB, post post) {
	id := ""      // ADD UUID
	created := "" //ADD time
	query, _ := db.Prepare("INSERT INTO post (id, title, content, user_id, community_id, created) VALUES (?, ?, ?, ?, ?)")
	query.Exec(id, post.title, post.content, post.user_id, post.community_id, created)
	defer query.Close()
}

func getPostsBySearchString(db *sql.DB, searchString string) []post {
	rows, err := db.Query("SELECT * FROM post WHERE title like '%" + searchString + "%' OR content like '%" + searchString + "%'")
	defer rows.Close()

	err = rows.Err()
	checkErr(err)

	postList := make([]post, 0)

	for rows.Next() {
		post := post{}
		err = rows.Scan(&post.id, &post.title, &post.content, &post.user_id, &post.community_id, &post.created)
		checkErr(err)

		postList = append(postList, post)
	}

	err = rows.Err()
	checkErr(err)

	return postList
}

func getPostsByUser(db *sql.DB, userID string) []post {
	rows, err := db.Query("SELECT * FROM post WHERE user_id='" + userID + "'")
	defer rows.Close()

	err = rows.Err()
	checkErr(err)

	postList := make([]post, 0)

	for rows.Next() {
		post := post{}
		err = rows.Scan(&post.id, &post.title, &post.content, &post.user_id, &post.community_id, &post.created)
		checkErr(err)

		postList = append(postList, post)
	}

	err = rows.Err()
	checkErr(err)

	return postList
}

func getPostsByCommunity(db *sql.DB, communityID string) []post {
	rows, err := db.Query("SELECT * FROM post WHERE community_id='" + communityID + "'")
	defer rows.Close()

	err = rows.Err()
	checkErr(err)

	postList := make([]post, 0)

	for rows.Next() {
		post := post{}
		err = rows.Scan(&post.id, &post.title, &post.content, &post.user_id, &post.community_id, &post.created)
		checkErr(err)

		postList = append(postList, post)
	}

	err = rows.Err()
	checkErr(err)

	return postList
}

func getPostById(db *sql.DB, id string) post {
	rows, _ := db.Query("SELECT * FROM post WHERE id = '" + id + "'")
	defer rows.Close()

	post := post{}

	for rows.Next() {
		rows.Scan(&post.id, &post.title, &post.content, &post.user_id, &post.community_id, &post.created)
	}

	return post
}

func updatePostInfo(db *sql.DB, post post) {
	stmt, err := db.Prepare("UPDATE post set title = ?, content = ?, user_id = ?, community_id = ?, created = ? where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(post.title, post.content, post.user_id, post.community_id, post.created, post.id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was affected")
	}
}

func deletePost(db *sql.DB, userID string) {
	stmt, err := db.Prepare("DELETE FROM user where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(userID)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 post was deleted")
	}
}
