package database

import (
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Comment struct {
	Id         int
	User_id    int
	Post_id    int
	Comment_id int
	Content    string
	Created    time.Time
}

func AddComment(comment Comment) {
	query, _ := DB.Prepare("INSERT INTO comment (user_id, post_id, comment_id, content, created) VALUES (?, NULLIF(?, 0), NULLIF(?, 0), ?, ?)")
	query.Exec(comment.User_id, comment.Post_id, comment.Comment_id, comment.Content, comment.Created)
	defer query.Close()
}

func GetCommentsByPost(postId int) []Comment {
	id := strconv.Itoa(postId)
	rows, err := DB.Query("SELECT * FROM comment WHERE post_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	commentList := make([]Comment, 0)

	for rows.Next() {
		comment := Comment{}
		err = rows.Scan(&comment.Id, &comment.User_id, &comment.Post_id, &comment.Comment_id, &comment.Content, &comment.Created)
		CheckErr(err)

		commentList = append(commentList, comment)
	}

	err = rows.Err()
	CheckErr(err)

	return commentList
}

func GetCommentsByUser(userId int) []Comment {
	id := strconv.Itoa(userId)
	rows, err := DB.Query("SELECT * FROM comment WHERE user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	commentList := make([]Comment, 0)

	for rows.Next() {
		comment := Comment{}
		err = rows.Scan(&comment.Id, &comment.User_id, &comment.Post_id, &comment.Comment_id, &comment.Content, &comment.Created)
		CheckErr(err)

		commentList = append(commentList, comment)
	}

	err = rows.Err()
	CheckErr(err)

	return commentList
}

func GetCommentsByComment(commentId int) []Comment {
	id := strconv.Itoa(commentId)
	rows, err := DB.Query("SELECT * FROM comment WHERE comment_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	commentList := make([]Comment, 0)

	for rows.Next() {
		comment := Comment{}
		err = rows.Scan(&comment.Id, &comment.User_id, &comment.Post_id, &comment.Comment_id, &comment.Content, &comment.Created)
		CheckErr(err)

		commentList = append(commentList, comment)
	}

	err = rows.Err()
	CheckErr(err)

	return commentList
}

func GetCommentById(id int) Comment {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT * FROM comment WHERE id = '" + id2 + "'")
	defer rows.Close()

	comment := Comment{}

	for rows.Next() {
		rows.Scan(&comment.Id, &comment.User_id, &comment.Post_id, &comment.Comment_id, &comment.Content, &comment.Created)
	}

	return comment
}

func UpdateCommentInfo(comment Comment) {
	query, err := DB.Prepare("UPDATE comment set user_id = ?, post_id = ?, comment_id = ?, content = ?, created = ? where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(comment.User_id, comment.Post_id, comment.Comment_id, comment.Content, comment.Created, comment.Id)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 comment was affected")
	}
}

func DeleteComment(commentId int) {
	query, err := DB.Prepare("DELETE FROM comment where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(commentId)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 comment was deleted")
	}
}
