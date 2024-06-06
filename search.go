package main

import (
	"ThreadCore/api"
	"ThreadCore/database"
	"html/template"
	"log"
	"net/http"
	"time"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search/" {
		http.Redirect(w, r, "/search/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/search.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	search := r.URL.Query().Get("q")
	media := r.URL.Query().Get("media") // media options  : posts, communities, comments , users
	sort := r.URL.Query().Get("sort")
	ChoosenTime := r.URL.Query().Get("time")

	var sortedPosts []database.PostInfo

	switch media {
	case "":
		searchedPost := database.GetPostsBySearchString(search)
		switch sort {
		case "popular":
			//TODO sort
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		case "new":
			api.NewestPost(searchedPost)
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		case "most_comment":
			//TODO sort
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		}

	case "posts":
		searchedPost := database.GetPostsBySearchString(search)
		switch sort {
		case "popular":
			//TODO sort
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		case "new":
			api.NewestPost(searchedPost)
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		case "most_comment":
			//TODO sort
			switch ChoosenTime {
			case "all_time":
				sortedPosts = searchedPost
			case "year":
				var YearTime = (time.Now().Add(-(time.Hour * 8764)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "month":
				var YearTime = (time.Now().Add(-(time.Hour * 744)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "week":
				var YearTime = (time.Now().Add(-(time.Hour * 168)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "day":
				var YearTime = (time.Now().Add(-(time.Hour * 24)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			case "hour":
				var YearTime = (time.Now().Add(-(time.Hour)))
				for i := 0; i < len(searchedPost); i++ {
					if !(searchedPost[i].Created.Before(YearTime)) {
						sortedPosts = append(sortedPosts, searchedPost[i])
					}
				}
			}
		}

	case "communities":
		searchedCommunities := database.GetCommunityBySearchString(search)

	case "users":
		searchedUser := database.GetUserBySearchString(search)

	}

	searchPage := struct {
		Search string
		Media  string
		Sort   string
		Time   string
		// Data   []string
	}{
		Search: search,
		Media:  media,
		Sort:   sort,
		Time:   ChoosenTime,
		// Data:   data,
	}

	err = tmpl.Execute(w, searchPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
