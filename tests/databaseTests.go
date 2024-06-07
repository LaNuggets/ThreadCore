package main

import (
	"ThreadCore/database"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	database.DB, err = sql.Open("sqlite3", "threadcore.db?_foreign_keys=on")
	database.CheckErr(err)

	// At the end of the program close the connnection
	defer database.DB.Close()

	u1 := database.User{Id: 1, Uuid: "uuid", Profile: "p", Banner: "b", Email: "email@email", Username: "username", Password: "password"}
	u2 := database.User{Id: 2, Uuid: "uuid2", Profile: "p2", Banner: "b2", Email: "email2", Username: "username2", Password: "password2"}
	u3 := database.User{Id: 3, Uuid: "uuid3", Profile: "p3", Banner: "b3", Email: "email3", Username: "username3", Password: "password3"}

	database.AddUser(u1)
	database.AddUser(u2)
	database.AddUser(u3)

	database.AddFriend(1, 2)
	database.AddFriend(2, 3)

	c1 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Minecraft", Description: "desc", User_id: 1}
	c2 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Fortnite", Description: "desc", User_id: 1}
	c3 := database.Community{Id: 0, Profile: "p", Banner: "b", Name: "Aurélien", Description: "desc", User_id: 2}
	database.AddCommunity(c1)
	database.AddCommunity(c2)
	database.AddCommunity(c3)

	p1 := database.Post{Id: 1, Title: "I like minecraft", Content: "minecraft is really cool and i like playing it a lot", User_id: 1, Community_id: 1, Created: (time.Now())}
	p2 := database.Post{Id: 2, Title: "I like my self", Content: "I am beautiful and evrybody need to now that", User_id: 2, Community_id: 2, Created: time.Now().Add(2 * time.Minute)}
	p3 := database.Post{Id: 3, Title: "Something about Minecraft", Content: "I want to talk about Minecraft", User_id: 1, Community_id: 1, Created: (time.Now()).Add(-(168 * time.Hour))}
	p4 := database.Post{Id: 4, Title: "Fortinite is shit", Content: "shit shit is", User_id: 3, Community_id: 2, Created: (time.Now()).Add(-(168 * time.Hour))}
	p5 := database.Post{Id: 5, Title: "Shit is Fortnite", Content: "is shit shit", User_id: 3, Community_id: 2, Created: (time.Now()).Add(-(744 * time.Hour))}

	database.AddPost(p1)
	database.AddPost(p2)
	database.AddPost(p3)
	database.AddPost(p4)
	database.AddPost(p5)

	l1 := database.Like{Id: 1, Rating: "like", Comment_id: 0, Post_id: 3, User_id: 1}
	l2 := database.Like{Id: 2, Rating: "like", Comment_id: 0, Post_id: 3, User_id: 2}
	l3 := database.Like{Id: 3, Rating: "like", Comment_id: 0, Post_id: 3, User_id: 3}
	l4 := database.Like{Id: 4, Rating: "like", Comment_id: 0, Post_id: 1, User_id: 1}
	l5 := database.Like{Id: 5, Rating: "like", Comment_id: 0, Post_id: 1, User_id: 2}
	l6 := database.Like{Id: 6, Rating: "like", Comment_id: 0, Post_id: 2, User_id: 1}

	database.AddLike(l1)
	database.AddLike(l2)
	database.AddLike(l3)
	database.AddLike(l4)
	database.AddLike(l5)
	database.AddLike(l6)

	database.AddUserCommunity(1, 2)
	database.AddUserCommunity(1, 3)
	database.AddUserCommunity(2, 3)
	database.AddUserCommunity(1, 1)

	comment1 := database.Comment{Id: 1, User_id: 1, Post_id: 1, Comment_id: 1, Content: "Yeah me to", Created: (time.Now()).Add(time.Minute)}
	comment2 := database.Comment{Id: 2, User_id: 2, Post_id: 2, Comment_id: 2, Content: "Wow fucking narcissistic", Created: (time.Now()).Add(2 * time.Minute)}
	comment3 := database.Comment{Id: 3, User_id: 3, Post_id: 1, Comment_id: 3, Content: "#Nerd", Created: time.Now().Add(-(168 * time.Hour))}
	comment4 := database.Comment{Id: 4, User_id: 2, Post_id: 3, Comment_id: 4, Content: "aglegle", Created: time.Now().Add(-(744 * time.Hour))}

	database.AddComment(comment1)
	database.AddComment(comment2)
	database.AddComment(comment3)
	database.AddComment(comment4)

	// fmt.Println(database.GetCommunitiesByNMembers()) // Still not working
	// fmt.Println("done")

}
