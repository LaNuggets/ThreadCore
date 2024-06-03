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
	database.DB, err = sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	database.CheckErr(err)

	// At the end of the program close the connnection
	defer database.DB.Close()

	u1 := database.User{Id: 0, Uuid: "uuid", Profile: "p", Banner: "b", Email: "email@email", Username: "username", Password: "password"}
	u2 := database.User{Id: 0, Uuid: "uuid2", Profile: "p2", Banner: "b2", Email: "email2", Username: "username2", Password: "password2"}
	database.AddUser(u1)
	database.AddUser(u2)

	c1 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Minecraft", Description: "its a community for the minecraft game", User_id: 1}
	c2 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Fortnite", Description: "its a community for the fortnite game", User_id: 1}
	c3 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Aurélien", Description: "welcome to the aurelien fan club", User_id: 2}
	database.AddCommunity(c1)
	database.AddCommunity(c2)
	database.AddCommunity(c3)

	p1 := database.Post{Id: 0, Title: "I like minecraft", Content: "minecraft is really cool and i like playing it a lot", User_id: 1, Community_id: 1, Created: time.Now()}
	p2 := database.Post{Id: 0, Title: "I like my self", Content: "I am beautiful and evrybody need to now that", User_id: 2, Community_id: 0, Created: time.Now()}
	database.AddPost(p1)
	database.AddPost(p2)

	database.AddUserCommunity(1, 2)
	database.AddUserCommunity(1, 3)
	database.AddUserCommunity(2, 3)

	fmt.Println(database.GetCommunitiesByNMembers()) // Still not working
	fmt.Println("done")

}
