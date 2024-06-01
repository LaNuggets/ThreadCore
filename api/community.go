package api

import (
	"ThreadCore/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// CREATE NEW COMMUNITY
func CreateCommunity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Check if user connected
	userUuid := CookieGetter("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for invalid user
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	community := database.GetCommunityByName(name)
	if (community != database.Community{}) {
		fmt.Println("Community already exists") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	r.ParseMultipartForm(10 << 20)

	// Get profile file or link from user
	profilePath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "link" {
		profilePath = r.FormValue("profileLink")
	} else {
		profile, handler1, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			profilePath = "/profileTemplate"
		} else {
			extension := strings.LastIndex(handler1.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext1 := handler1.Filename[extension:] //obtain the extension in ext variable
			profilePath = "communities/profile/" + name + ext1
			GetFileFromForm(profile, handler1, err, profilePath)
		}
	}

	// Get banner file or link from user
	bannerPath := ""

	bannerOption := r.FormValue("bannerOption")
	if bannerOption == "link" {
		bannerPath = r.FormValue("bannerLink")
	} else {
		banner, handler2, err := r.FormFile("banner")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			bannerPath = "/bannerTemplate.png"
		} else {
			extension := strings.LastIndex(handler2.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext2 := handler2.Filename[extension:] //obtain the extension in ext variable
			bannerPath = "communities/banner/" + name + ext2
			GetFileFromForm(banner, handler2, err, bannerPath)
		}
	}

	community = database.Community{Id: 0, Profile: profilePath, Banner: bannerPath, Name: name, Following: 0, User_id: user.Id}
	database.AddCommunity(community)

	http.Redirect(w, r, "/community/"+name, http.StatusSeeOther)
}

// UPDATE EXISTING COMMUNITY
func UpdateCommunity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	id := r.FormValue("communityId")
	community := database.GetCommunityById(id)
	if (community == database.Community{}) {
		fmt.Println("community does not exist") // TO-DO : send error community not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := CookieGetter("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for invalid user
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	} else if community.User_id != user.Id {
		fmt.Println("user not author of community") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	nameOption := r.FormValue("nameOption")

	newName := ""
	if nameOption == "change" {
		newName = r.FormValue("newName")
		checkCommunity := database.GetCommunityByName(newName)
		if (checkCommunity != database.Community{}) {
			fmt.Println("Community already exists") // TO-DO : Send error message for invalid name
			http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
			return
		}
	} else {
		newName = community.Name
	}

	checkCommunity := database.GetCommunityByName(newName)
	if (checkCommunity != database.Community{}) {
		fmt.Println("Community already exists") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	r.ParseMultipartForm(10 << 20)

	// Get profile file or link from user
	profilePath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "remove" {
		profilePath = "/profileTemplate.png"
	} else if profileOption == "keep" {
		profilePath = community.Profile
	} else if profileOption == "link" {
		profilePath = r.FormValue("profileLink")
	} else {
		profile, handler1, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			profilePath = "/profileTemplate.png"
		} else {
			extension := strings.LastIndex(handler1.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext1 := handler1.Filename[extension:] //obtain the extension in ext variable
			profilePath = "communities/profile/" + strconv.Itoa(community.Id) + ext1
			GetFileFromForm(profile, handler1, err, profilePath)
		}
	}

	// Get profile file or link from user
	bannerPath := ""

	bannerOption := r.FormValue("bannerOption")
	if bannerOption == "remove" {
		bannerPath = "/bannerTemplate.png"
	} else if bannerOption == "keep" {
		bannerPath = community.Banner
	} else if bannerOption == "link" {
		bannerPath = r.FormValue("bannerLink")
	} else {
		banner, handler2, err := r.FormFile("banner")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			bannerPath = "/bannerTemplate.png"
		} else {
			extension := strings.LastIndex(handler2.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext2 := handler2.Filename[extension:] //obtain the extension in ext variable
			bannerPath = "communities/banner/" + strconv.Itoa(community.Id) + ext2
			GetFileFromForm(banner, handler2, err, bannerPath)
		}
	}

	community = database.Community{Id: community.Id, Profile: profilePath, Banner: bannerPath, Name: newName, Following: 0, User_id: user.Id}
	database.UpdateCommunityInfo(community)

	http.Redirect(w, r, "/community/"+newName, http.StatusSeeOther)
}
