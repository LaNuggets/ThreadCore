package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Community(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/community/" {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/community.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	communityName := strings.ReplaceAll(r.URL.Path, "/community/", "")
	if strings.Contains(communityName, "/") {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
	}

	//ProfilePicture and connection check
	userUuid := api.GetCookie("uuid", r)
	user := database.GetUserByUuid(userUuid, w, r)

	community := database.GetCommunityByName(communityName, w, r)

	//Following info
	followcount := len(database.GetUsersByCommunity(community.Id, w, r))
	isFollowing := database.ExistsUserCommunity(user.Id, community.Id)

	type UserPageInfo struct {
		Connected bool
		Profile   string
		Username  string
		Uuid      string
		Id        int
	}
	userInfo := UserPageInfo{
		Connected: userUuid != "",
		Profile:   user.Profile,
		Username:  api.GetCookie("username", r),
		Uuid:      user.Uuid,
		Id:        user.Id,
	}

	communityPage := struct {
		User            UserPageInfo
		Community       database.Community
		CommunityExists bool
		FollowCount     int
		IsFollowing     bool
	}{
		User:            userInfo,
		Community:       community,
		CommunityExists: community != database.Community{},
		FollowCount:     followcount,
		IsFollowing:     isFollowing,
	}

	err = tmpl.Execute(w, communityPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
