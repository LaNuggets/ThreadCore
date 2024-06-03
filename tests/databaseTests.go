package main

import (
	"ThreadCore/database"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	database.DB, err = sql.Open("sqlite3", "threadcore.db")
	database.CheckErr(err)

	// At the end of the program close the connnection
	defer database.DB.Close()

	c1 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Minecraft", Following: 0}
	c2 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Fortnite", Following: 0}
	c3 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Aurélien", Following: 0}
	database.AddCommunity(c1)
	database.AddCommunity(c2)
	database.AddCommunity(c3)

	u1 := database.User{Id: 0, Uuid: "uuid", Profile: "p", Banner: "b", Email: "email", Username: "username", Password: "password"}
	u2 := database.User{Id: 0, Uuid: "uuid2", Profile: "p2", Banner: "b2", Email: "email2", Username: "username2", Password: "password2"}
	database.AddUser(u1)
	database.AddUser(u2)

	database.AddUserCommunity(1, 2)
	database.AddUserCommunity(2, 2)
	database.AddUserCommunity(2, 3)

	fmt.Println(database.GetCommunitiesByNMembers())
	fmt.Println("done")

	p1 := database.Post{Id: 0, Title: "I like minecraft", Content: "minecraft is really cool and i like playing it a lot.", User_id: 1, Community_id: 1, Created: time.Now()}
	p2 := database.Post{Id: 1, Title: "I like my self", Content: "I am beautiful and evrybody need to now that.", User_id: 2, Community_id: 2, Created: time.Now()}
	p3 := database.Post{Id: 3, Title: "New glitch", Content: "Yesturday i just found a new bug in minecraft whiche allow player to duplicate everything.", User_id: 2, Community_id: 1, Created: time.Now()}

	database.AddPost(p1)
	database.AddPost(p2)
	database.AddPost(p3)

	comment1 := database.Comment{Id: 0, User_id: 2, Post_id: 3, Comment_id: 0, Content: "Prety cool things", Created: time.Now()}
	database.AddComment(comment1)
}
