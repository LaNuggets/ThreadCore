package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to database
	db, err := sql.Open("sqlite3", "../threadcore.db")
	fmt.Println("create DB:")
	checkErr(err)
	// defer close
	defer db.Close()

	dropTables := `
DROP TABLE IF EXISTS user; 
DROP TABLE IF EXISTS message; 
DROP TABLE IF EXISTS post; 
DROP TABLE IF EXISTS comment; 
DROP TABLE IF EXISTS like; 
DROP TABLE IF EXISTS community; 
DROP TABLE IF EXISTS user_community; 
DROP TABLE IF EXISTS groupchat; 
DROP TABLE IF EXISTS user_groupchat; 
DROP TABLE IF EXISTS friend; 
DROP TABLE IF EXISTS friend_request; 
	`
	_, err = db.Exec(dropTables)
	fmt.Println("drop table:")
	checkErr(err)

	createTables := `
CREATE TABLE user(id INTEGER PRIMARY KEY, profilePicture TEXT, email TEXT, username TEXT, password TEXT); 
CREATE TABLE message(id INTEGER PRIMARY KEY, user_id INTEGER, groupchat_id INTEGER, friend_id INTEGER, message INTEGER, sent TIMESTAMP, FOREIGN KEY (user_id) REFERENCES friend(user_id), FOREIGN KEY (groupchat_id) REFERENCES groupchat(id), FOREIGN KEY (friend_id) REFERENCES friend(friend_id)); 
CREATE TABLE post(id INTEGER PRIMARY KEY, title TEXT, content TEXT, user_id INTEGER, community_id INTEGER, created TIMESTAMP, FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (community_id) REFERENCES community(id)); 
CREATE TABLE comment(id INTEGER PRIMARY KEY, user_id INTEGER, post_id INTEGER, comment_id INTEGER, content TEXT, created TIMESTAMP, FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (post_id) REFERENCES post(id), FOREIGN KEY (comment_id) REFERENCES comment(id)); 
CREATE TABLE like(id INTEGER PRIMARY KEY, rating INTEGER, comment_id INTEGER, post_id INTEGER, user_id INTEGER, FOREIGN KEY (comment_id) REFERENCES comment(id), FOREIGN KEY (post_id) REFERENCES post(id), FOREIGN KEY (user_id) REFERENCES user(id)); 
CREATE TABLE community(id INTEGER PRIMARY KEY, name TEXT); 
CREATE TABLE user_community(user_id INTEGER, community_id INTEGER, PRIMARY KEY(user_id, community_id), FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (community_id) REFERENCES community(id)); 
CREATE TABLE groupchat(id INTEGER PRIMARY KEY, name TEXT); 
CREATE TABLE user_groupchat(user_id INTEGER, groupchat_id INTEGER, PRIMARY KEY(user_id, groupchat_id), FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (groupchat_id) REFERENCES groupchat(id)); 
CREATE TABLE friend(user_id INTEGER, friend_id INTEGER, PRIMARY KEY(user_id, friend_id), FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (friend_id) REFERENCES user(id)); 
CREATE TABLE friend_request(user_id INTEGER, request_id INTEGER, PRIMARY KEY(user_id, request_id), FOREIGN KEY (user_id) REFERENCES user(id),FOREIGN KEY (request_id) REFERENCES user(id)); 
	`
	_, err = db.Exec(createTables)
	fmt.Println("create table:")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
