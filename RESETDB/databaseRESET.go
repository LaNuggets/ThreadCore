package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to database
	db, err := sql.Open("sqlite3", "./threadcore.db")
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
CREATE TABLE user(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, uuid VARCHAR(255),  profile VARCHAR(255), banner VARCHAR(255), email VARCHAR(64), username VARCHAR(20), password VARCHAR(255));
CREATE TABLE message(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user_id INTEGER, groupchat_id INTEGER, friend_id INTEGER, message VARCHAR(255), sent TIMESTAMP, FOREIGN KEY (user_id) REFERENCES friend(user_id), FOREIGN KEY (groupchat_id) REFERENCES groupchat(id), FOREIGN KEY (friend_id) REFERENCES friend(friend_id));
CREATE TABLE post(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, title VARCHAR(32), content VARCHAR(255), media VARCHAR(255),user_id INTEGER, community_id INTEGER, created TIMESTAMP, FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (community_id) REFERENCES community(id));
CREATE TABLE comment(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, user_id INTEGER, post_id INTEGER, comment_id INTEGER, content VARCHAR(255), created TIMESTAMP, FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (post_id) REFERENCES post(id), FOREIGN KEY (comment_id) REFERENCES comment(id));
CREATE TABLE like(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, rating INTEGER, comment_id INTEGER, post_id INTEGER, user_id INTEGER, FOREIGN KEY (comment_id) REFERENCES comment(id), FOREIGN KEY (post_id) REFERENCES post(id), FOREIGN KEY (user_id) REFERENCES user(id));
CREATE TABLE community(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,  profile VARCHAR(255), banner VARCHAR(255), name VARCHAR(32), following INTEGER NOT NULL);
CREATE TABLE user_community(user_id INTEGER, community_id INTEGER, PRIMARY KEY(user_id, community_id), FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (community_id) REFERENCES community(id));
CREATE TABLE groupchat(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name VARCHAR(32));
CREATE TABLE user_groupchat(user_id INTEGER, groupchat_id INTEGER, PRIMARY KEY(user_id, groupchat_id), FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (groupchat_id) REFERENCES groupchat(id));
CREATE TABLE friend(user_id INTEGER, friend_id INTEGER, PRIMARY KEY(user_id, friend_id), FOREIGN KEY (user_id) REFERENCES user(id), FOREIGN KEY (friend_id) REFERENCES user(id));
CREATE TABLE friend_request(user_id INTEGER, request_id INTEGER, PRIMARY KEY(user_id, request_id), FOREIGN KEY (user_id) REFERENCES user(id),FOREIGN KEY (request_id) REFERENCES user(id));
	`
	_, err = db.Exec(createTables)
	fmt.Println("create table:")
	checkErr(err)

	fmt.Println("Successfuly created the database!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
