package api

import (
	"ThreadCore/database"
	"sort"
	"time"
)

type test []database.PostInfo

func (a test) Len() int {
	return len(a)
}
func (a test) Less(i, j int) bool { return a[j].Created.Before(a[i].Created) }
func (a test) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func NewestPost(posts []database.PostInfo) {
	sort.Sort(test(posts))
}

func YearAgo(posts []database.PostInfo) []database.PostInfo {
	var YearTime = (time.Now().Add(-(time.Hour * 8764)))
	var sortedPosts []database.PostInfo
	for i := 0; i < len(posts); i++ {
		if !(posts[i].Created.Before(YearTime)) {
			sortedPosts = append(sortedPosts, posts[i])
		}
	}
	return sortedPosts
}

func MonthAgo(posts []database.PostInfo) []database.PostInfo {
	var MonthTime = (time.Now().Add(-(time.Hour * 744)))
	var sortedPosts []database.PostInfo
	for i := 0; i < len(posts); i++ {
		if !(posts[i].Created.Before(MonthTime)) {
			sortedPosts = append(sortedPosts, posts[i])
		}
	}
	return sortedPosts
}

func WeekAgo(posts []database.PostInfo) []database.PostInfo {
	var weekTime = (time.Now().Add(-(time.Hour * 168)))
	var sortedPosts []database.PostInfo
	for i := 0; i < len(posts); i++ {
		if !(posts[i].Created.Before(weekTime)) {
			sortedPosts = append(sortedPosts, posts[i])
		}
	}
	return sortedPosts
}

func DayAgo(posts []database.PostInfo) []database.PostInfo {
	var dayTime = (time.Now().Add(-(time.Hour * 24)))
	var sortedPosts []database.PostInfo
	for i := 0; i < len(posts); i++ {
		if !(posts[i].Created.Before(dayTime)) {
			sortedPosts = append(sortedPosts, posts[i])
		}
	}
	return sortedPosts
}

func HourAgo(posts []database.PostInfo) []database.PostInfo {
	var hourTime = (time.Now().Add(-(time.Hour * 1)))
	var sortedPosts []database.PostInfo
	for i := 0; i < len(posts); i++ {
		if !(posts[i].Created.Before(hourTime)) {
			sortedPosts = append(sortedPosts, posts[i])
		}
	}
	return sortedPosts
}
