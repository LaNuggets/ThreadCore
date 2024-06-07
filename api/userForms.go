package api

import (
	"ThreadCore/database"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// SIGNUP
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// check valid username and email
	username := r.FormValue("username")
	checkUsername := database.GetUserByUsername(username)
	if (checkUsername != database.User{}) {
		fmt.Println("username taken") // TO-DO : send error user not found
		http.Redirect(w, r, "/connection", http.StatusSeeOther)
		return
	}
	email := r.FormValue("email")
	checkEmail := database.GetUserByEmail(email)
	if (checkEmail != database.User{}) {
		fmt.Println("email taken") // TO-DO : send error user not found
		http.Redirect(w, r, "/connection", http.StatusSeeOther)
		return
	}
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("passwordConfirm")
	if passwordConfirm != password {
		fmt.Println("password and passwordConfirm dont match") // TO-DO : Send error message for confirm password
		http.Redirect(w, r, "/connection", http.StatusSeeOther)
		return
	} else if password == "" {
		fmt.Println("password is null") // TO-DO : Send error message for input password
		http.Redirect(w, r, "/connection", http.StatusSeeOther)
		return
	}
	password = HashPassword(password)
	uuid := GetNewUuid()

	user := database.User{Id: 0, Uuid: uuid, Profile: "/static/images/profileTemplate.png", Banner: "/static/images/bannerTemplate.png", Email: email, Username: username, Password: password}
	database.AddUser(user)

	SetCookie("uuid", user.Uuid, w)
	SetCookie("username", user.Username, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// LOGIN
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")
	user := database.GetUserByUsername(username)
	if (user == database.User{}) {
		fmt.Println("username not found") // TO-DO : send error user not found
		http.Redirect(w, r, "/connection", http.StatusSeeOther)
		return
	}
	email := r.FormValue("email")
	user2 := database.GetUserByEmail(email)
	if (user2 == database.User{}) {
		fmt.Println("email not found") // TO-DO : send error user not found
		http.Redirect(w, r, "/connection", http.StatusSeeOther)
		return
	}
	if reflect.DeepEqual(user, user2) {
		fmt.Println("user not found check username or email") // TO-DO : send error user not found
		http.Redirect(w, r, "/connection", http.StatusSeeOther)
		return
	}
	if !CheckPasswordHash(user.Password, user.Password) {
		fmt.Println("Wrong password") // TO-DO : Send error message for wrong password
		http.Redirect(w, r, "/connection", http.StatusSeeOther)
		return
	}

	SetCookie("uuid", user.Uuid, w)
	SetCookie("username", user.Username, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DISCONNECT
func Disconnect(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	DeleteCookie("uuid", w)
	DeleteCookie("username", w)
	http.Redirect(w, r, "/connection", http.StatusSeeOther)
}

// UPDATE EXISTING user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	uuidToUpdate := r.FormValue("uuid")
	userToUpdate := database.GetUserByUuid(uuidToUpdate)
	if (userToUpdate == database.User{}) {
		fmt.Println("user does not exist") // TO-DO : send error user not found
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
		return
	} else if userToUpdate.Id != user.Id {
		fmt.Println("user not author of user") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
		return
	}

	// check valid username and email
	username := r.FormValue("username")
	checkUsername := database.GetUserByUsername(username)
	if (checkUsername != database.User{}) {
		fmt.Println("username taken") // TO-DO : send error user not found
		http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
		return
	}
	email := r.FormValue("email")
	checkEmail := database.GetUserByEmail(email)
	if (checkEmail != database.User{}) {
		fmt.Println("email taken") // TO-DO : send error user not found
		http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
		return
	}

	password := ""
	passwordOption := r.FormValue("passwordOption")
	if passwordOption == "change" {
		oldPassword := r.FormValue("oldPassword")
		password = r.FormValue("password")
		passwordConfirm := r.FormValue("passwordConfirm")

		if !CheckPasswordHash(oldPassword, user.Password) {
			fmt.Println("Wrong password") // TO-DO : Send error message for wrong password
			http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
			return
		} else if passwordConfirm != password {
			fmt.Println("password and passwordConfirm dont match") // TO-DO : Send error message for confirm password
			http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
			return
		} else if password == "" {
			fmt.Println("password is null") // TO-DO : Send error message for input password
			http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
			return
		}

		password = HashPassword(password)
	} else {
		password = user.Password
	}

	r.ParseMultipartForm(10 << 20)

	// Get profile file or link from user
	profilePath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "remove" {
		profilePath = "/static/images/profileTemplate.png"
	} else if profileOption == "keep" {
		profilePath = user.Profile
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
				profilePath = "/static/images/users/profile/" + user.Uuid + e
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
		bannerPath = user.Banner
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
				bannerPath = "/static/images/users/banner/" + user.Uuid + ext
				GetFileFromForm(banner, handler, err, bannerPath)
			} else {
				fmt.Println("The file is  not in an image format")
				return //if not an image format
			}
		}
	}

	user = database.User{Id: user.Id, Uuid: user.Uuid, Profile: profilePath, Banner: bannerPath, Email: email, Username: username, Password: password}
	database.UpdateUserInfo(user)

	http.Redirect(w, r, "/user/"+userToUpdate.Username, http.StatusSeeOther)
}

// DELETE user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	uuidToDelete := r.FormValue("uuid")
	userToDelete := database.GetUserByUuid(uuidToDelete)
	if (userToDelete == database.User{}) {
		fmt.Println("user does not exist") // TO-DO : send error user not found
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/user/"+userToDelete.Username, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/user/"+userToDelete.Username, http.StatusSeeOther)
		return
	} else if userToDelete.Id != user.Id {
		fmt.Println("user not author of user") // TO-DO : Send error message for user not allowed action
		http.Redirect(w, r, "/user/"+userToDelete.Username, http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion") // TO-DO : Send error message need to confirm before submiting
		http.Redirect(w, r, "/user/"+user.Username, http.StatusSeeOther)
		return
	} else {
		database.DeleteUser(user.Id)
	}

	//Send confirmation message
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	uuidToFollow := r.FormValue("uuid")
	userToFollow := database.GetUserByUuid(uuidToFollow)
	if (userToFollow == database.User{}) {
		fmt.Println("community does not exist") // TO-DO : send error community not found
		http.Redirect(w, r, "/user/"+userToFollow.Username, http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/user/"+userToFollow.Username, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/user/"+userToFollow.Username, http.StatusSeeOther)
		return
	}

	if database.ExistsFriend(user.Id, userToFollow.Id) {
		fmt.Println("user already following this user")
	} else {
		database.AddFriend(user.Id, userToFollow.Id)
	}
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	uuidToFollow := r.FormValue("uuid")
	userToFollow := database.GetUserByUuid(uuidToFollow)
	if (userToFollow == database.User{}) {
		fmt.Println("community does not exist") // TO-DO : send error community not found
		http.Redirect(w, r, "/user/"+userToFollow.Username, http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie") // TO-DO : Send error message for user not connected
		http.Redirect(w, r, "/user/"+userToFollow.Username, http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid)
	if (user == database.User{}) {
		fmt.Println("user not found") // TO-DO : Send error message for user not found
		http.Redirect(w, r, "/user/"+userToFollow.Username, http.StatusSeeOther)
		return
	}

	if database.ExistsFriend(user.Id, userToFollow.Id) {
		database.DeleteFriend(user.Id, userToFollow.Id)
	} else {
		fmt.Println("user already not following this user")
	}
}
