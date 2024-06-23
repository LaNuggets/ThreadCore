package api

import (
	"ThreadCore/database"
	"fmt"
	"net/http"
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
	checkUsername := database.GetUserByUsername(username, w, r)
	if (checkUsername != database.User{}) {
		fmt.Println("username taken")
		http.Redirect(w, r, "/?type=error&message=Username+taken%2C+please+choose+another+one+!", http.StatusSeeOther)
		return
	}
	email := r.FormValue("email")
	checkEmail := database.GetUserByEmail(email, w, r)
	if (checkEmail != database.User{}) {
		fmt.Println("email taken")
		http.Redirect(w, r, "/?type=error&message=Email+taken%2C+please+choose+another+one+!", http.StatusSeeOther)
		return
	}
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("passwordConfirm")
	if passwordConfirm != password {
		fmt.Println("password and passwordConfirm dont match")
		http.Redirect(w, r, "/?type=error&message=Password+confiramtion+is+wrong%2C+please+try+again+!", http.StatusSeeOther)
		return
	} else if password == "" {
		fmt.Println("password is null")
		http.Redirect(w, r, "/?type=error&message=Password+is+empty+!", http.StatusSeeOther)
		return
	}
	password = HashPassword(password, w, r)
	uuid := GetNewUuid()

	user := database.User{Id: 0, Uuid: uuid, Profile: "/static/images/profileTemplate.png", Banner: "/static/images/bannerTemplate.png", Email: email, Username: username, Password: password}
	database.AddUser(user, w, r)

	SetCookie("uuid", user.Uuid, w)
	http.Redirect(w, r, "/?type=success&message=Account+successfully+created+!", http.StatusSeeOther)
}

// LOGIN
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")
	user := database.GetUserByUsername(username, w, r)
	if (user == database.User{}) {
		fmt.Println("username not found")
		http.Redirect(w, r, "/?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	}
	email := r.FormValue("email")
	user2 := database.GetUserByEmail(email, w, r)
	if (user2 == database.User{}) {
		fmt.Println("email not found")
		http.Redirect(w, r, "/?type=error&message=Email+not+found+!", http.StatusSeeOther)
		return
	}
	if user.Uuid != user2.Uuid {
		fmt.Println("user not found check username or email")
		http.Redirect(w, r, "/?type=error&message=Username+or+Email+invalid+!", http.StatusSeeOther)
		return
	}
	password := r.FormValue("password")
	if !CheckPasswordHash(password, user.Password) {
		fmt.Println("Wrong password")
		http.Redirect(w, r, "/?type=error&message=Wrong+Password+!", http.StatusSeeOther)
		return
	}

	SetCookie("uuid", user.Uuid, w)
	http.Redirect(w, r, "/?type=success&message=Connection+successful+!", http.StatusSeeOther)
}

// DISCONNECT
func Disconnect(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	DeleteCookie("uuid", w)
	http.Redirect(w, r, "/?type=success&message=Successfully+Disconnected+!", http.StatusSeeOther)
}

// UPDATE EXISTING user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	uuidToUpdate := r.FormValue("uuid")
	userToUpdate := database.GetUserByUuid(uuidToUpdate, w, r)
	if (userToUpdate == database.User{}) {
		fmt.Println("user does not exist")
		http.Redirect(w, r, "/?type=error&message=User+does+not+exist+!", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie")
		http.Redirect(w, r, "/user/"+userToUpdate.Username+"?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found")
		http.Redirect(w, r, "/user/"+userToUpdate.Username+"?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	} else if userToUpdate.Id != user.Id {
		fmt.Println("user not author of user")
		http.Redirect(w, r, "/user/"+userToUpdate.Username+"?type=error&message=You+are+not+the+owner+of+this+account+!", http.StatusSeeOther)
		return
	}

	// check valid username and email
	username := r.FormValue("username")
	checkUsername := database.GetUserByUsername(username, w, r)
	if (checkUsername != database.User{} && checkUsername.Uuid != userUuid) {
		fmt.Println("username taken")
		http.Redirect(w, r, "/user/"+userToUpdate.Username+"?type=error&message=Username+taken+!", http.StatusSeeOther)
		return
	}
	email := r.FormValue("email")
	checkEmail := database.GetUserByEmail(email, w, r)
	if (checkEmail != database.User{} && checkEmail.Uuid != userUuid) {
		fmt.Println("email taken")
		http.Redirect(w, r, "/user/"+userToUpdate.Username+"?type=error&message=Email+taken+!", http.StatusSeeOther)
		return
	}

	oldPassword := r.FormValue("oldPassword")
	password := r.FormValue("newPassword")
	passwordConfirm := r.FormValue("newPasswordConfirm")

	if oldPassword == "" {
		password = user.Password
	} else {
		if !CheckPasswordHash(oldPassword, user.Password) {
			fmt.Println("Wrong password")
			http.Redirect(w, r, "/user/"+userToUpdate.Username+"?type=error&message=Wrong+Password+!", http.StatusSeeOther)
			return
		} else if passwordConfirm != password {
			fmt.Println("password and passwordConfirm dont match")
			http.Redirect(w, r, "/user/"+userToUpdate.Username+"?type=error&message=Wrong+Password+confirmation+!", http.StatusSeeOther)
			return
		} else if password == "" {
			fmt.Println("password is null")
			http.Redirect(w, r, "/user/"+userToUpdate.Username+"?type=error&message=New+password+empty+!", http.StatusSeeOther)
			return
		} else {
			password = HashPassword(password, w, r)
		}
	}

	r.ParseMultipartForm(10 << 20)

	// Get profile file or link from user
	profilePath := ""

	profileOption := r.FormValue("profileOption")
	if profileOption == "remove" {
		DeleteFile(userToUpdate.Profile)
		profilePath = "/static/images/profileTemplate.png"
	} else if profileOption == "keep" {
		profilePath = user.Profile
	} else if profileOption == "link" {
		DeleteFile(userToUpdate.Profile)
		profilePath = r.FormValue("profileLink")
	} else {
		DeleteFile(userToUpdate.Profile)
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
				profilePath = "/static/images/users/profiles/" + user.Uuid + e
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
		DeleteFile(userToUpdate.Banner)
		bannerPath = "/static/images/bannerTemplate.png"
	} else if bannerOption == "keep" {
		bannerPath = user.Banner
	} else if bannerOption == "link" {
		DeleteFile(userToUpdate.Banner)
		bannerPath = r.FormValue("bannerLink")
	} else {
		DeleteFile(userToUpdate.Banner)
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
				bannerPath = "/static/images/users/banners/" + user.Uuid + ext
				GetFileFromForm(banner, handler, err, bannerPath)
			} else {
				fmt.Println("The file is  not in an image format")
				return //if not an image format
			}
		}
	}

	user = database.User{Id: user.Id, Uuid: user.Uuid, Profile: profilePath, Banner: bannerPath, Email: email, Username: username, Password: password}
	database.UpdateUserInfo(user, w, r)

	http.Redirect(w, r, "/user/"+username+"?type=success&message=Account+successfully+updated+!", http.StatusSeeOther)
}

// DELETE user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	uuidToDelete := r.FormValue("uuid")
	userToDelete := database.GetUserByUuid(uuidToDelete, w, r)
	if (userToDelete == database.User{}) {
		fmt.Println("user to delete does not exist")
		http.Redirect(w, r, "/?type=error&message=User+does+not+exist+!", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie")
		http.Redirect(w, r, "/user/"+userToDelete.Username+"?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found")
		http.Redirect(w, r, "/user/"+userToDelete.Username+"?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	} else if userToDelete.Id != user.Id {
		fmt.Println("user not author of user")
		http.Redirect(w, r, "/user/"+userToDelete.Username+"?type=error&message=User+not+allowed+to+do+this+action+!", http.StatusSeeOther)
		return
	}

	confirm := r.FormValue("confirm")
	password := r.FormValue("password")
	if confirm != "true" {
		fmt.Println("user did not confirm deletion")
		http.Redirect(w, r, "/user/"+user.Username+"?type=error&message=Confirm+deletion+!", http.StatusSeeOther)
		return
	} else {
		if !CheckPasswordHash(password, userToDelete.Password) {
			fmt.Println("Wrong Password")
			http.Redirect(w, r, "/user/"+user.Username+"?type=error&message=Wrong+password+!", http.StatusSeeOther)
			return
		} else {
			DeleteFile(userToDelete.Profile)
			DeleteFile(userToDelete.Banner)
			database.DeleteUser(user.Id, w, r)
		}
	}

	DeleteCookie("uuid", w)
	http.Redirect(w, r, "/?type=success&message=User+successfully+deleted+!", http.StatusSeeOther)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	action := r.FormValue("action")

	uuidToFollow := r.FormValue("uuid")
	userToFollow := database.GetUserByUuid(uuidToFollow, w, r)
	if (userToFollow == database.User{}) {
		fmt.Println("community does not exist")
		http.Redirect(w, r, action+"?type=error&message=Community+not+found+!", http.StatusSeeOther)
		return
	}

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie")
		http.Redirect(w, r, action+"?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found")
		http.Redirect(w, r, action+"?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	}

	if database.ExistsFriend(user.Id, userToFollow.Id, w, r) {
		fmt.Println("user already following this user")
		http.Redirect(w, r, action+"?type=error&message=You+are+already+following+"+userToFollow.Username+"+!", http.StatusSeeOther)
	} else {
		database.AddFriend(user.Id, userToFollow.Id, w, r)
		http.Redirect(w, r, action+"?type=success&message=You+are+now+following+"+userToFollow.Username+"+!", http.StatusSeeOther)
	}

}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	uuidToFollow := r.FormValue("uuid")
	userToFollow := database.GetUserByUuid(uuidToFollow, w, r)
	if (userToFollow == database.User{}) {
		fmt.Println("community does not exist")
		http.Redirect(w, r, "/user/"+userToFollow.Username+"?type=error&message=Community+not+found+!", http.StatusSeeOther)
		return
	}
	action := r.FormValue("action")

	// Check if user connected and allowed to modify
	userUuid := GetCookie("uuid", r)
	if userUuid == "" {
		fmt.Println("no uuid found in cookie")
		http.Redirect(w, r, action+"?type=error&message=User+not+connected+!", http.StatusSeeOther)
		return
	}
	user := database.GetUserByUuid(userUuid, w, r)
	if (user == database.User{}) {
		fmt.Println("user not found")
		http.Redirect(w, r, action+"?type=error&message=User+not+found+!", http.StatusSeeOther)
		return
	}

	if database.ExistsFriend(user.Id, userToFollow.Id, w, r) {
		database.DeleteFriend(user.Id, userToFollow.Id, w, r)
		http.Redirect(w, r, action+"?type=success&message=You+are+not+following+"+userToFollow.Username+"+anymore+!", http.StatusSeeOther)
	} else {
		fmt.Println("user already not following this user")
		http.Redirect(w, r, action+"?type=error&message=Your+are+not+following+"+userToFollow.Username+"+!", http.StatusSeeOther)
	}
}
