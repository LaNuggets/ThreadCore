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

type CommentInfo struct {
	Id         int
	User_id    int
	Username   string
	Post_id    int
	Comment_id int
	Content    string
	Created    time.Time
}

type TempCommentInfo struct {
	Id         int
	User_id    int
	Username   string
	Post_id    *int
	Comment_id *int
	Content    string
	Created    time.Time
}

func AddComment(comment Comment) {
	query, _ := DB.Prepare("INSERT INTO comment (user_id, post_id, comment_id, content, created) VALUES (?, NULLIF(?, 0), NULLIF(?, 0), ?, ?)")
	query.Exec(comment.User_id, comment.Post_id, comment.Comment_id, comment.Content, comment.Created)
	defer query.Close()
}

// func GetCommentsBySearchString(searchString string) []CommentInfo {
// 	rows, err := DB.Query("SELECT comment.id, comment.user_id, user.username, post.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id WHERE comment.content LIKE '%" + searchString + "%'")
// 	defer rows.Close()

// 	err = rows.Err()
// 	CheckErr(err)

// 	commentList := make([]CommentInfo, 0)

// 	for rows.Next() {
// 		tempCommentInfo := TempCommentInfo{}
// 		err = rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.User_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
// 		CheckErr(err)
// 		commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
// 		if tempCommentInfo.Comment_id != nil {
// 			commentInfo.Comment_id = *tempCommentInfo.Comment_id
// 		}
// 		if tempCommentInfo.Post_id != nil {
// 			commentInfo.Post_id = *tempCommentInfo.Post_id
// 		}
// 		commentList = append(commentList, commentInfo)
// 	}

// 	err = rows.Err()
// 	CheckErr(err)

// 	return commentList
// }

func GetCommentsByPost(postId int) []CommentInfo {
	id := strconv.Itoa(postId)
	rows, err := DB.Query("SELECT comment.id, comment.user_id, user.username, post.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id FROM comment WHERE post_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	commentList := make([]CommentInfo, 0)

	for rows.Next() {
		tempCommentInfo := TempCommentInfo{}
		err = rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.User_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
		CheckErr(err)
		commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
		if tempCommentInfo.Comment_id != nil {
			commentInfo.Comment_id = *tempCommentInfo.Comment_id
		}
		if tempCommentInfo.Post_id != nil {
			commentInfo.Post_id = *tempCommentInfo.Post_id
		}
		commentList = append(commentList, commentInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return commentList
}

func GetCommentsByUser(userId int) []CommentInfo {
	id := strconv.Itoa(userId)
	rows, err := DB.Query("SELECT comment.id, comment.user_id, user.username, post.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id FROM comment WHERE user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	commentList := make([]CommentInfo, 0)

	for rows.Next() {
		tempCommentInfo := TempCommentInfo{}
		err = rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.User_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
		CheckErr(err)
		commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
		if tempCommentInfo.Comment_id != nil {
			commentInfo.Comment_id = *tempCommentInfo.Comment_id
		}
		if tempCommentInfo.Post_id != nil {
			commentInfo.Post_id = *tempCommentInfo.Post_id
		}
		commentList = append(commentList, commentInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return commentList
}

func GetCommentsByComment(commentId int) []CommentInfo {
	id := strconv.Itoa(commentId)
	rows, err := DB.Query("SELECT comment.id, comment.user_id, user.username, post.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id FROM comment WHERE comment_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err)

	commentList := make([]CommentInfo, 0)

	for rows.Next() {
		tempCommentInfo := TempCommentInfo{}
		err = rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.User_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
		CheckErr(err)
		commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
		if tempCommentInfo.Comment_id != nil {
			commentInfo.Comment_id = *tempCommentInfo.Comment_id
		}
		if tempCommentInfo.Post_id != nil {
			commentInfo.Post_id = *tempCommentInfo.Post_id
		}
		commentList = append(commentList, commentInfo)
	}

	err = rows.Err()
	CheckErr(err)

	return commentList
}

func GetCommentById(id int) CommentInfo {
	id2 := strconv.Itoa(id)
	rows, _ := DB.Query("SELECT comment.id, comment.user_id, user.username, post.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id FROM comment WHERE id = '" + id2 + "'")
	defer rows.Close()

	tempCommentInfo := TempCommentInfo{}

	for rows.Next() {
		rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.User_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
	}

	commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
	if tempCommentInfo.Comment_id != nil {
		commentInfo.Comment_id = *tempCommentInfo.Comment_id
	}
	if tempCommentInfo.Post_id != nil {
		commentInfo.Post_id = *tempCommentInfo.Post_id
	}

	return commentInfo
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
