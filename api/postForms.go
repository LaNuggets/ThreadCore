package api

import (
	"ThreadCore/database"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CREATE NEW Post
func CreatePost(w http.ResponseWriter, r *http.Request) {
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

	postId := r.FormValue("postId")
	title := r.FormValue("title")
	content := r.FormValue("content")
	id := r.FormValue("communityId")
	communityId, _ := strconv.Atoi(id)

	r.ParseMultipartForm(10 << 20)

	// Get profile file or link from user
	mediaPath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "link" {
		mediaPath = r.FormValue("profileLink")
	} else {
		profile, handler1, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			mediaPath = "/mediaTemplate"
		} else {
			extension := strings.LastIndex(handler1.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext1 := handler1.Filename[extension:] //obtain the extension in ext variable
			mediaPath = "/static/images/posts/" + postId + ext1
			GetFileFromForm(profile, handler1, err, mediaPath)
		}
	}

	post := database.Post{Id: 0, Title: title, Content: content, Media: mediaPath, User_id: user.Id, Community_id: communityId, Created: time.Now()}
	database.AddPost(post)

	http.Redirect(w, r, "/post/"+postId, http.StatusSeeOther)
}

// UPDATE EXISTING Post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	postId := r.FormValue("postId")
	id, _ := strconv.Atoi(postId)
	post := database.GetPostById(id)
	if (post == database.Post{}) {
		fmt.Println("post does not exist") // TO-DO : send error post not found
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := CookieGetter("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/post/"+strconv.Itoa(post.Id), http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
		return
	} else if post.User_id != user.Id {
		fmt.Println("user not author of post") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
		return
	}

	nameOption := r.FormValue("nameOption")

	newName := ""
	if nameOption == "change" {
		newName = r.FormValue("newName")
		checkPost := database.GetPostByName(newName)
		if (checkPost != database.Post{}) {
			fmt.Println("Post already exists") // TO-DO : Send error message for invalid name
			http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
			return
		}
	} else {
		newName = post.Name
	}

	checkPost := database.GetPostByName(newName)
	if (checkPost != database.Post{}) {
		fmt.Println("Post already exists") // TO-DO : Send error message for invalid name
		http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
		return
	}

	r.ParseMultipartForm(10 << 20)

	// Get profile file or link from user
	mediaPath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "remove" {
		mediaPath = "/profileTemplate.png"
	} else if profileOption == "keep" {
		mediaPath = post.Profile
	} else if profileOption == "link" {
		mediaPath = r.FormValue("profileLink")
	} else {
		profile, handler1, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			mediaPath = "/profileTemplate.png"
		} else {
			extension := strings.LastIndex(handler1.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext1 := handler1.Filename[extension:] //obtain the extension in ext variable
			mediaPath = "/static/images/posts/" + strconv.Itoa(post.Id) + ext1
			GetFileFromForm(profile, handler1, err, mediaPath)
		}
	}

	// Get profile file or link from user
	bannerPath := ""

	bannerOption := r.FormValue("bannerOption")
	if bannerOption == "remove" {
		bannerPath = "/bannerTemplate.png"
	} else if bannerOption == "keep" {
		bannerPath = post.Banner
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
			bannerPath = "/static/images/communities/banner/" + strconv.Itoa(post.Id) + ext2
			GetFileFromForm(banner, handler2, err, bannerPath)
		}
	}

	post = database.Post{Id: post.Id, Profile: mediaPath, Banner: bannerPath, Name: newName, Following: 0, User_id: user.Id}
	database.UpdatePostInfo(post)

	http.Redirect(w, r, "/post/"+newName, http.StatusSeeOther)
}

// DELETE Post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	id := r.FormValue("PostId")
	post := database.GetPostById(id)
	if (post == database.Post{}) {
		fmt.Println("post does not exist") // TO-DO : send error post not found
		http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := CookieGetter("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
		return
	} else if post.User_id != user.Id {
		fmt.Println("user not author of post") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/post/"+post.Name, http.StatusSeeOther)
		return
	} else {
		database.DeletePost(post.Id)
	}

	//Send confirmation message
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
