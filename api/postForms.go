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
	userUuid := GetCookie("uuid", r)
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

	postUuid := GetNewUuid()
	title := r.FormValue("title")
	content := r.FormValue("content")
	id := r.FormValue("communityId")
	communityId, _ := strconv.Atoi(id)

	r.ParseMultipartForm(10 << 20)

	// Get image or video file or link from user
	mediaPath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "link" {
		mediaPath = r.FormValue("profileLink")
	} else {
		profile, handler, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			mediaPath = "/static/images/mediaTemplate.png"
		} else {
			extension := strings.LastIndex(handler.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext := handler.Filename[extension:] //obtain the extension in ext variable
			e := strings.ToLower(ext)
			if e == ".png" || e == ".jpeg" || e == ".jpg" || e == ".gif" || e == ".svg" || e == ".avif" || e == ".apng" || e == ".webp" || e == ".mp4" || e == ".webm" || e == ".ogg" {
				mediaPath = "/static/images/posts/" + postUuid + ext
				GetFileFromForm(profile, handler, err, mediaPath)
			} else {
				fmt.Println("The file is  not in an image or video format")
				return //if not an image or video format
			}
		}
	}

	post := database.Post{Id: 0, Uuid: postUuid, Title: title, Content: content, Media: mediaPath, User_id: user.Id, Community_id: communityId, Created: time.Now()}
	database.AddPost(post)

	http.Redirect(w, r, "/post/"+postUuid, http.StatusSeeOther)
}

// UPDATE EXISTING Post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	postUuid := r.FormValue("postUuid")
	id, _ := strconv.Atoi(postUuid)
	post := database.GetPostById(id)
	if (post == database.PostInfo{}) {
		fmt.Println("post does not exist") // TO-DO : send error post not found
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/post/"+postUuid, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/post/"+postUuid, http.StatusSeeOther)
		return
	} else if post.User_id != user.Id {
		fmt.Println("user not author of post") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/post/"+postUuid, http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	communityid := r.FormValue("communityId")
	communityId, _ := strconv.Atoi(communityid)

	r.ParseMultipartForm(10 << 20)

	// Get image or video file or link from user
	mediaPath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "remove" {
		mediaPath = "/static/images/mediaTemplate.png"
	} else if profileOption == "keep" {
		mediaPath = post.Media
	} else if profileOption == "link" {
		mediaPath = r.FormValue("profileLink")
	} else {
		profile, handler, err := r.FormFile("profile")

		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
			mediaPath = "/static/images/mediaTemplate.png"
		} else {
			extension := strings.LastIndex(handler.Filename, ".") //obtain the extension after the dot
			if extension == -1 {
				fmt.Println("The file has no extension")
				return //if no extension is present print failure
			}
			ext := handler.Filename[extension:] //obtain the extension in ext variable
			e := strings.ToLower(ext)
			if e == ".png" || e == ".jpeg" || e == ".jpg" || e == ".gif" || e == ".svg" || e == ".avif" || e == ".apng" || e == ".webp" || e == ".mp4" || e == ".webm" || e == ".ogg" {
				mediaPath = "/static/images/posts/" + postUuid + ext
				GetFileFromForm(profile, handler, err, mediaPath)
			} else {
				fmt.Println("The file is  not in an image or video format")
				return //if not an image or video format
			}
		}
	}

	updatedPost := database.Post{Id: 0, Title: title, Content: content, Media: mediaPath, User_id: user.Id, Community_id: communityId, Created: post.Created}
	database.UpdatePostInfo(updatedPost)

	http.Redirect(w, r, "/post/"+postUuid, http.StatusSeeOther)
}

// DELETE Post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	postId := r.FormValue("PostId")
	id, _ := strconv.Atoi(postId)
	post := database.GetPostById(id)
	if (post == database.PostInfo{}) {
		fmt.Println("post does not exist") // TO-DO : send error post not found
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("Uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/post/"+postId, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/post/"+postId, http.StatusSeeOther)
		return
	} else if post.User_id != user.Id {
		fmt.Println("user not author of post") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/post/"+postId, http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/post/"+postId, http.StatusSeeOther)
		return
	} else {
		database.DeletePost(post.Id)
	}

	//Send confirmation message
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
