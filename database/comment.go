package database

import (
	"database/sql"
	"log"
	"net/http"
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
	Profile    string
	Post_id    int
	Comment_id int
	Content    string
	Created    time.Time
	Time       string
}

type TempCommentInfo struct {
	Id         int
	User_id    int
	Username   string
	Profile    string
	Post_id    *int
	Comment_id *int
	Content    string
	Created    time.Time
}

/*
!AddComment function open data base and add comment by using the INSERT INTO sql command she take as argument an comment type and a writer and request.
*/
func AddComment(comment Comment, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, _ := db.Prepare("INSERT INTO comment (user_id, post_id, comment_id, content, created) VALUES (?, NULLIF(?, 0), NULLIF(?, 0), ?, ?)")
	query.Exec(comment.User_id, comment.Post_id, comment.Comment_id, comment.Content, comment.Created)
	defer query.Close()

}

// func GetCommentsBySearchString(searchString string) []CommentInfo {
// 	rows, err := db.Query("SELECT comment.id, comment.user_id, user.username, post.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id WHERE comment.content LIKE '%" + searchString + "%'")
// 	defer rows.Close()

// 	err = rows.Err()
// 	CheckErr(err, w, r)

// 	commentList := make([]CommentInfo, 0)

// 	for rows.Next() {
// 		tempCommentInfo := TempCommentInfo{}
// 		err = rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.User_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
// 		CheckErr(err, w, r)
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
// 	CheckErr(err, w, r)

// 	return commentList
// }
/*
!GetCommentsByPost function open data base and get comments by post by using the SELECT sql command she take as argument an int type and a writer and request and return a slice of CommentInfo.
*/
func GetCommentsByPost(postId int, w http.ResponseWriter, r *http.Request) []CommentInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(postId)
	rows, err := db.Query("SELECT comment.id, comment.user_id, user.username, user.profile, comment.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id WHERE post_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	commentList := make([]CommentInfo, 0)

	for rows.Next() {
		tempCommentInfo := TempCommentInfo{}
		err = rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.Profile, &tempCommentInfo.Post_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
		CheckErr(err, w, r)
		commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Profile: tempCommentInfo.Profile, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
		if tempCommentInfo.Comment_id != nil {
			commentInfo.Comment_id = *tempCommentInfo.Comment_id
		}
		if tempCommentInfo.Post_id != nil {
			commentInfo.Post_id = *tempCommentInfo.Post_id
		}
		commentList = append(commentList, commentInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return commentList
}

/*
!GetCommentsByUser function open data base and get comments by user by using the SELECT sql command she take as argument an int type and a writer and request and return a slice of CommentInfo.
*/
func GetCommentsByUser(userId int, w http.ResponseWriter, r *http.Request) []CommentInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(userId)
	rows, err := db.Query("SELECT comment.id, comment.user_id, user.username, user.profile, comment.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id WHERE user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	commentList := make([]CommentInfo, 0)

	for rows.Next() {
		tempCommentInfo := TempCommentInfo{}
		err = rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.Profile, &tempCommentInfo.Post_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
		CheckErr(err, w, r)
		commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Profile: tempCommentInfo.Profile, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
		if tempCommentInfo.Comment_id != nil {
			commentInfo.Comment_id = *tempCommentInfo.Comment_id
		}
		if tempCommentInfo.Post_id != nil {
			commentInfo.Post_id = *tempCommentInfo.Post_id
		}
		commentList = append(commentList, commentInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return commentList
}

/*
!GetCommentsByComment function open data base and get comments by comment by using the SELECT sql command she take as argument an int type and a writer and request and return a slice of CommentInfo.
*/
func GetCommentsByComment(commentId int, w http.ResponseWriter, r *http.Request) []CommentInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(commentId)
	rows, err := db.Query("SELECT comment.id, comment.user_id, user.username, user.profile, comment.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id WHERE comment_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	commentList := make([]CommentInfo, 0)

	for rows.Next() {
		tempCommentInfo := TempCommentInfo{}
		err = rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.Profile, &tempCommentInfo.Post_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
		CheckErr(err, w, r)
		commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Profile: tempCommentInfo.Profile, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
		if tempCommentInfo.Comment_id != nil {
			commentInfo.Comment_id = *tempCommentInfo.Comment_id
		}
		if tempCommentInfo.Post_id != nil {
			commentInfo.Post_id = *tempCommentInfo.Post_id
		}
		commentList = append(commentList, commentInfo)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return commentList
}

/*
!GetCommentById function open data base and get comments by id by using the SELECT sql command she take as argument an int type and a writer and request and return a slice of CommentInfo.
*/
func GetCommentById(id int, w http.ResponseWriter, r *http.Request) CommentInfo {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id2 := strconv.Itoa(id)
	rows, _ := db.Query("SELECT comment.id, comment.user_id, user.username, user.profile, comment.post_id, comment.comment_id, comment.content, comment.created FROM comment INNER JOIN user ON user.id = comment.user_id WHERE id = '" + id2 + "'")
	defer rows.Close()

	tempCommentInfo := TempCommentInfo{}

	for rows.Next() {
		rows.Scan(&tempCommentInfo.Id, &tempCommentInfo.User_id, &tempCommentInfo.Username, &tempCommentInfo.Profile, &tempCommentInfo.Post_id, &tempCommentInfo.Comment_id, &tempCommentInfo.Content, &tempCommentInfo.Created)
	}

	commentInfo := CommentInfo{Id: tempCommentInfo.Id, User_id: tempCommentInfo.User_id, Username: tempCommentInfo.Username, Profile: tempCommentInfo.Profile, Post_id: 0, Comment_id: 0, Content: tempCommentInfo.Content, Created: tempCommentInfo.Created}
	if tempCommentInfo.Comment_id != nil {
		commentInfo.Comment_id = *tempCommentInfo.Comment_id
	}
	if tempCommentInfo.Post_id != nil {
		commentInfo.Post_id = *tempCommentInfo.Post_id
	}

	return commentInfo
}

/*
!UpdateCommentInfo function open data base and update comment information by using the UPDATE sql command she take as argument an Comment type and a writer and request.
*/
func UpdateCommentInfo(comment Comment, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE comment set user_id = ?, post_id = NULLIF(?, 0), comment_id = NULLIF(?, 0), content = ?, created = ? where id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(comment.User_id, comment.Post_id, comment.Comment_id, comment.Content, comment.Created, comment.Id)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 comment was affected")
	}
}

/*
!DeleteComment function open data base and delete comment by using the DELETE sql command she take as argument an int type and a writer and request.
*/
func DeleteComment(commentId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM comment where id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(commentId)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 comment was deleted")
	}
}
