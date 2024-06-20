package database

import (
	"database/sql"
	"log"
	"net/http"
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

type TempLike struct {
	Id         int
	Rating     string
	Comment_id *int
	Post_id    *int
	User_id    int
}

func AddLike(like Like, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, _ := db.Prepare("INSERT INTO like (rating, comment_id, post_id, user_id) VALUES (?, NULLIF(?, 0), NULLIF(?, 0), ?)")
	query.Exec(like.Rating, like.Comment_id, like.Post_id, like.User_id)
	defer query.Close()
}

func GetLikesByPost(postId int, w http.ResponseWriter, r *http.Request) []Like {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(postId)
	rows, err := db.Query("SELECT * FROM like WHERE post_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	likeList := make([]Like, 0)

	for rows.Next() {
		templike := TempLike{}
		err = rows.Scan(&templike.Id, &templike.Rating, &templike.Comment_id, &templike.Post_id, &templike.User_id)
		CheckErr(err, w, r)
		like := Like{Id: templike.Id, Rating: templike.Rating, Comment_id: 0, Post_id: 0, User_id: templike.User_id}
		if templike.Post_id != nil {
			like.Post_id = *templike.Post_id
		} else if templike.Comment_id != nil {
			like.Comment_id = *templike.Comment_id
		}
		likeList = append(likeList, like)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return likeList
}

func GetLikesByUser(userId int, w http.ResponseWriter, r *http.Request) []Like {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(userId)
	rows, err := db.Query("SELECT * FROM like WHERE user_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	likeList := make([]Like, 0)

	for rows.Next() {
		templike := TempLike{}
		err = rows.Scan(&templike.Id, &templike.Rating, &templike.Comment_id, &templike.Post_id, &templike.User_id)
		CheckErr(err, w, r)
		like := Like{Id: templike.Id, Rating: templike.Rating, Comment_id: 0, Post_id: 0, User_id: templike.User_id}
		if templike.Post_id != nil {
			like.Post_id = *templike.Post_id
		} else if templike.Comment_id != nil {
			like.Comment_id = *templike.Comment_id
		}
		likeList = append(likeList, like)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return likeList
}

func GetLikesByComment(commentId int, w http.ResponseWriter, r *http.Request) []Like {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id := strconv.Itoa(commentId)
	rows, err := db.Query("SELECT * FROM like WHERE comment_id='" + id + "'")
	defer rows.Close()

	err = rows.Err()
	CheckErr(err, w, r)

	likeList := make([]Like, 0)

	for rows.Next() {
		templike := TempLike{}
		err = rows.Scan(&templike.Id, &templike.Rating, &templike.Comment_id, &templike.Post_id, &templike.User_id)
		CheckErr(err, w, r)
		like := Like{Id: templike.Id, Rating: templike.Rating, Comment_id: 0, Post_id: 0, User_id: templike.User_id}
		if templike.Post_id != nil {
			like.Post_id = *templike.Post_id
		} else if templike.Comment_id != nil {
			like.Comment_id = *templike.Comment_id
		}
		likeList = append(likeList, like)
	}

	err = rows.Err()
	CheckErr(err, w, r)

	return likeList
}

func GetLikeById(id int, w http.ResponseWriter, r *http.Request) Like {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	id2 := strconv.Itoa(id)
	rows, _ := db.Query("SELECT * FROM like WHERE id = '" + id2 + "'")
	defer rows.Close()

	templike := TempLike{}

	for rows.Next() {
		rows.Scan(&templike.Id, &templike.Rating, &templike.Comment_id, &templike.Post_id, &templike.User_id)
	}

	like := Like{Id: templike.Id, Rating: templike.Rating, Comment_id: 0, Post_id: 0, User_id: templike.User_id}
	if templike.Post_id != nil {
		like.Post_id = *templike.Post_id
	} else if templike.Comment_id != nil {
		like.Comment_id = *templike.Comment_id
	}
	return like
}

func GetLikeByUserComment(user_id int, comment_id int, w http.ResponseWriter, r *http.Request) Like {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	userId := strconv.Itoa(user_id)
	commentId := strconv.Itoa(comment_id)
	rows, _ := db.Query("SELECT * FROM like WHERE user_id = '" + userId + "' AND comment_id = '" + commentId + "'")
	defer rows.Close()

	templike := TempLike{}

	for rows.Next() {
		rows.Scan(&templike.Id, &templike.Rating, &templike.Comment_id, &templike.Post_id, &templike.User_id)
	}

	like := Like{Id: templike.Id, Rating: templike.Rating, Comment_id: 0, Post_id: 0, User_id: templike.User_id}
	if templike.Post_id != nil {
		like.Post_id = *templike.Post_id
	} else if templike.Comment_id != nil {
		like.Comment_id = *templike.Comment_id
	}
	return like
}

func GetLikeByUserPost(user_id int, post_id int, w http.ResponseWriter, r *http.Request) Like {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	userId := strconv.Itoa(user_id)
	postId := strconv.Itoa(post_id)
	rows, _ := db.Query("SELECT * FROM like WHERE user_id = '" + userId + "' AND post_id = '" + postId + "'")
	defer rows.Close()

	templike := TempLike{}

	for rows.Next() {
		rows.Scan(&templike.Id, &templike.Rating, &templike.Comment_id, &templike.Post_id, &templike.User_id)
	}

	like := Like{Id: templike.Id, Rating: templike.Rating, Comment_id: 0, Post_id: 0, User_id: templike.User_id}
	if templike.Post_id != nil {
		like.Post_id = *templike.Post_id
	} else if templike.Comment_id != nil {
		like.Comment_id = *templike.Comment_id
	}
	return like
}

func UpdateLike(like Like, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("UPDATE like SET rating = ? where id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(like.Rating, like.Id)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 user was affected")
	}
}

func DeleteLike(likeId int, w http.ResponseWriter, r *http.Request) {
	//Open the database connection
	db, err := sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	CheckErr(err, w, r)
	// Close the batabase at the end of the program
	defer db.Close()

	query, err := db.Prepare("DELETE FROM like WHERE id = ?")
	CheckErr(err, w, r)
	defer query.Close()

	res, err := query.Exec(likeId)
	CheckErr(err, w, r)

	affected, err := res.RowsAffected()
	CheckErr(err, w, r)

	if affected > 1 {
		log.Fatal("Error : More than 1 like was deleted")
	}
}
