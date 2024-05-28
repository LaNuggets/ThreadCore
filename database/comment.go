package database

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Comment struct {
	Id         *int
	User_id    int
	Post_id    int
	Comment_id int
	Content    string
	Created    time.Time
}

func AddComment(comment Comment) {
	query, _ := DB.Prepare("INSERT INTO comment (user_id, comment_id, comment_id, content, created) VALUES (?, ?, ?, ?, ?)")
	query.Exec(comment.User_id, comment.Post_id, comment.Comment_id, comment.Content, comment.Created)
	defer query.Close()
}

func GetCommentsByPost(postId string) []Comment {
	rows, err := DB.Query("SELECT * FROM comment WHERE post_id='" + postId + "'")
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

func GetCommentsByUser(userID string) []Comment {
	rows, err := DB.Query("SELECT * FROM comment WHERE user_id='" + userID + "'")
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

func GetCommentsByComment(commentID string) []Comment {
	rows, err := DB.Query("SELECT * FROM comment WHERE comment_id='" + commentID + "'")
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

func GetCommentById(id string) Comment {
	rows, _ := DB.Query("SELECT * FROM comment WHERE id = '" + id + "'")
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

func DeleteComment(commentID string) {
	query, err := DB.Prepare("DELETE FROM comment where id = ?")
	CheckErr(err)
	defer query.Close()

	res, err := query.Exec(commentID)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	if affected > 1 {
		log.Fatal("Error : More than 1 comment was deleted")
	}
}
