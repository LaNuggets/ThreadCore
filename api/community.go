package api

import (
	"ThreadCore/database"
	"fmt"
	"net/http"
	"strings"
)

func CreateCommunity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	author := CookieGetter("Uuid")
	if len(author) == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for invalid user
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
	} else {
		user := database.g
	}

	name := r.FormValue("name")
	community := database.GetCommunityByName(name)
	if (community != database.Community{}) {
		fmt.Println("Community already exists") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
	}

	r.ParseMultipartForm(10 << 20)

	profile, handler1, err := r.FormFile("profile")
	profilePath := ""
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
		fmt.Println("The extension of", handler1.Filename, "is", ext1)
		profilePath = "communities/profile/" + name + ext1
		GetFileFromForm(profile, handler1, err, profilePath)
	}

	banner, handler2, err := r.FormFile("banner")
	bannerPath := ""
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
		fmt.Println("The extension of", handler2.Filename, "is", ext2)

		bannerPath = "communities/banner/" + name + ext2
		GetFileFromForm(banner, handler2, err, bannerPath)
	}

	community = database.Community{Id: 0, Profile: profilePath, Banner: bannerPath, Name: name, Following: 0}
	database.AddCommunity(community)

	http.Redirect(w, r, "/community/"+name, http.StatusSeeOther)
}

func UpdateCommunity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	name := r.FormValue("name")
	community := database.GetCommunityByName(name)
	if (community == database.Community{}) {
		fmt.Println("Community does not exist")
		http.Redirect(w, r, "/community/"+name, http.StatusSeeOther)
	}

	newName := r.FormValue("newName")
	checkCommunity := database.GetCommunityByName(newName)
	if (checkCommunity != database.Community{}) {
		fmt.Println("Community already exists") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/community/"+name, http.StatusSeeOther)
	}

	r.ParseMultipartForm(10 << 20)

	profile, handler1, err := r.FormFile("profile")
	profilePath := ""
	if err == http.ErrMissingFile {
		fmt.Println("no file uploaded")
		profilePath = community.Profile
	} else {
		extension := strings.LastIndex(handler1.Filename, ".") //obtain the extension after the dot
		if extension == -1 {
			fmt.Println("The file has no extension")
			return //if no extension is present print failure
		}
		ext1 := handler1.Filename[extension:] //obtain the extension in ext variable
		fmt.Println("The extension of", handler1.Filename, "is", ext1)
		profilePath = "communities/profile/" + name + ext1
		GetFileFromForm(profile, handler1, err, profilePath)
	}

	banner, handler2, err := r.FormFile("banner")
	bannerPath := ""
	if err == http.ErrMissingFile {
		fmt.Println("no file uploaded")
		bannerPath = community.Banner
	} else {
		extension := strings.LastIndex(handler2.Filename, ".") //obtain the extension after the dot
		if extension == -1 {
			fmt.Println("The file has no extension")
			return //if no extension is present print failure
		}
		ext2 := handler2.Filename[extension:] //obtain the extension in ext variable
		fmt.Println("The extension of", handler2.Filename, "is", ext2)

		bannerPath = "communities/banner/" + name + ext2
		GetFileFromForm(banner, handler2, err, bannerPath)
	}

	id := community.Id
	community = database.Community{Id: id, Profile: profilePath, Banner: bannerPath, Name: name, Following: 0}
	database.AddCommunity(community)

	http.Redirect(w, r, "/community/"+name, http.StatusSeeOther)
}
