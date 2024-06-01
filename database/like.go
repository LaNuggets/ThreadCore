package database

import (
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Like struct {
	Id         int
	Rating     string
	Comment_id int
	Post_id    int
	User_id    int
}

func AddLike(like Like) {
	query, _ := DB.Prepare("INSERT INTO like (rating, comment_id, post_id, user_id) VALUES (?, NULLIF(?, 0), NULLIF(?, 0), ?)")
	query.Exec(like.Rating, like.Comment_id, like.Post_id, like.User_id)
	defer query.Close()
}

func GetLikesByPost(postId int) []Like {
	id := strconv.Itoa(postId)
	rows, err := DB.Query("SELECT * FROM like WHERE post_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	likeList := make([]Like, 0)

	for rows.Next() {
		like := Like{}
		err = rows.Scan(&like.Id, &like.Rating, &like.Comment_id, &like.Post_id, &like.User_id)
		CheckErr(err)

		likeList = append(likeList, like)
	}

	err = rows.Err()
	CheckErr(err)

	return likeList
}

func GetLikesByUser(userId int) []Like {
	id := strconv.Itoa(userId)
	rows, err := DB.Query("SELECT * FROM like WHERE user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	likeList := make([]Like, 0)

	for rows.Next() {
		like := Like{}
		err = rows.Scan(&like.Id, &like.Rating, &like.Comment_id, &like.Post_id, &like.User_id)
		CheckErr(err)

		likeList = append(likeList, like)
	}

	err = rows.Err()
	CheckErr(err)

	return likeList
}

func GetLikesByComment(commentId int) []Like {
	id := strconv.Itoa(commentId)
	rows, err := DB.Query("SELECT * FROM like WHERE comment_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	likeList := make([]Like, 0)

	for rows.Next() {
		like := Like{}
		err = rows.Scan(&like.Id, &like.Rating, &like.Comment_id, &like.Post_id, &like.User_id)
		CheckErr(err)

		likeList = append(likeList, like)
	}

	err = rows.Err()
	CheckErr(err)

	return likeList
}

func GetLikeById(id int) Like {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT * FROM comment WHERE id = '" + id2 + "'")
	defer rows.Close()

	like := Like{}

	for rows.Next() {
		rows.Scan(&like.Id, &like.Rating, &like.Comment_id, &like.Post_id, &like.User_id)
	}

	return like
}

func UpdateLike(like Like) {
	query, err := DB.Prepare("UPDATE like set rating = ?, comment_id = ?, post_id = ?, user_id = ? where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(like.Rating, like.Comment_id, like.Post_id, like.User_id, like.Id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was affected")
	}
}

func DeleteLike(likeId int) {
	query, err := DB.Prepare("DELETE FROM like where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(likeId)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}
