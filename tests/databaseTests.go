package main

import (
	"ThreadCore/database"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	database.DB, err = sql.Open("sqlite3", "threadcore.db")
	database.CheckErr(err)

	// At the end of the program close the connnection
	defer database.DB.Close()

	c1 := database.Community{Id: nil, Name: "Minecraft"}
	c2 := database.Community{Id: nil, Name: "Fortnite"}
	c3 := database.Community{Id: nil, Name: "Aurélien"}
	database.AddCommunity(c1)
	database.AddCommunity(c2)
	database.AddCommunity(c3)

	//u1 := database.User{Id: 0, Uuid: "uuid", ProfilePicture: "pp", Email: "email", Username: "username", Password: "password"}

	database.DeleteCommunity(database.GetCommunityById(1).Id)
	fmt.Println("deleted")

	fmt.Println(database.GetCommunitiesByNMembers())
	fmt.Println("done")
}
