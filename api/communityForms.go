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
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
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
		profile, handler, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			profilePath = "/static/images/profileTemplate.png"
		} else {
			extension := strings.LastIndex(handler.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext := handler.Filename[extension:] //obtain the extension in ext variable
			e := strings.ToLower(ext)
			if e == ".png" || e == ".jpeg" || e == ".jpg" || e == ".gif" || e == ".svg" || e == ".avif" || e == ".apng" || e == ".webp" {
				profilePath = "/static/images/communities/profile/" + name + ext
				GetFileFromForm(profile, handler, err, profilePath)
			} else {
				fmt.Println("The file is  not in an image format")
				return //if not an image format
			}
		}
	}

	// Get banner file or link from user
	bannerPath := ""

	bannerOption := r.FormValue("bannerOption")
	if bannerOption == "link" {
		bannerPath = r.FormValue("bannerLink")
	} else {
		banner, handler, err := r.FormFile("banner")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			bannerPath = "/static/images/bannerTemplate.png"
		} else {
			extension := strings.LastIndex(handler.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext := handler.Filename[extension:] //obtain the extension in ext variable
			e := strings.ToLower(ext)
			if e == ".png" || e == ".jpeg" || e == ".jpg" || e == ".gif" || e == ".svg" || e == ".avif" || e == ".apng" || e == ".webp" {
				bannerPath = "/static/images/communities/banner/" + name + ext
				GetFileFromForm(banner, handler, err, bannerPath)
			} else {
				fmt.Println("The file is  not in an image format")
				return //if not an image format
			}
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
	id2, _ := strconv.Atoi(id)
	community := database.GetCommunityById(id2)
	if (community == database.Community{}) {
		fmt.Println("community does not exist") // TO-DO : send error community not found
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := CookieGetter("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	} else if community.User_id != user.Id {
		fmt.Println("user not author of community") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	newName := r.FormValue("newName")

	r.ParseMultipartForm(10 << 20)

	// Get profile file or link from user
	profilePath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "remove" {
		profilePath = "/static/images/profileTemplate.png.png"
	} else if profileOption == "keep" {
		profilePath = community.Profile
	} else if profileOption == "link" {
		profilePath = r.FormValue("profileLink")
	} else {
		profile, handler, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			profilePath = "/static/images/profileTemplate.png"
		} else {
			extension := strings.LastIndex(handler.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext := handler.Filename[extension:] //obtain the extension in ext variable
			e := strings.ToLower(ext)
			if e == ".png" || e == ".jpeg" || e == ".jpg" || e == ".gif" || e == ".svg" || e == ".avif" || e == ".apng" || e == ".webp" {
				profilePath = "/static/images/communities/profile/" + strconv.Itoa(community.Id) + ext
				GetFileFromForm(profile, handler, err, profilePath)
			} else {
				fmt.Println("The file is  not in an image format")
				return //if not an image format
			}
		}
	}

	// Get profile file or link from user
	bannerPath := ""

	bannerOption := r.FormValue("bannerOption")
	if bannerOption == "remove" {
		bannerPath = "/static/images/bannerTemplate.png"
	} else if bannerOption == "keep" {
		bannerPath = community.Banner
	} else if bannerOption == "link" {
		bannerPath = r.FormValue("bannerLink")
	} else {
		banner, handler, err := r.FormFile("banner")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			bannerPath = "/static/images/bannerTemplate.png"
		} else {
			extension := strings.LastIndex(handler.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext := handler.Filename[extension:] //obtain the extension in ext variable
			e := strings.ToLower(ext)
			if e == ".png" || e == ".jpeg" || e == ".jpg" || e == ".gif" || e == ".svg" || e == ".avif" || e == ".apng" || e == ".webp" {
				bannerPath = "/static/images/communities/banner/" + strconv.Itoa(community.Id) + ext
				GetFileFromForm(banner, handler, err, bannerPath)
			} else {
				fmt.Println("The file is  not in an image format")
				return //if not an image format
			}
		}
	}

	community = database.Community{Id: community.Id, Profile: profilePath, Banner: bannerPath, Name: newName, Following: 0, User_id: user.Id}
	database.UpdateCommunityInfo(community)

	http.Redirect(w, r, "/community/"+newName, http.StatusSeeOther)
}

// DELETE COMMUNITY
func DeleteCommunity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	id := r.FormValue("communityId")
	id2, _ := strconv.Atoi(id)
	community := database.GetCommunityById(id2)
	if (community == database.Community{}) {
		fmt.Println("community does not exist") // TO-DO : send error community not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := CookieGetter("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	} else if community.User_id != user.Id {
		fmt.Println("user not author of community") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	} else {
		database.DeleteCommunity(community.Id)
	}

	//Send confirmation message
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func FollowCommunity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	communityId := r.FormValue("communityId")
	communityid, _ := strconv.Atoi(communityId)
	community := database.GetCommunityById(communityid)
	if (community == database.Community{}) {
		fmt.Println("community does not exist") // TO-DO : send error community not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := CookieGetter("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	if database.ExistsUserCommunity(user.Id, communityid) {
		fmt.Println("user already following this community")
	} else {
		database.AddUserCommunity(user.Id, communityid)
	}
}

func UnfollowCommunity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	communityId := r.FormValue("communityId")
	communityid, _ := strconv.Atoi(communityId)
	community := database.GetCommunityById(communityid)
	if (community == database.Community{}) {
		fmt.Println("community does not exist") // TO-DO : send error community not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := CookieGetter("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/community/"+community.Name, http.StatusSeeOther)
		return
	}

	if database.ExistsUserCommunity(user.Id, communityid) {
		database.DeleteUserCommunity(user.Id, communityid)
	} else {
		fmt.Println("user already not following this community")
	}
}
