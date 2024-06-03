package api

import (
	"ThreadCore/database"
	"sort"
)

// func MostPopularPost(posts []database.PostDisplay) []database.PostDisplay {

// }

type test []database.PostDisplay

func (a test) Len() int {
	return len(a)
}
func (a test) Less(i, j int) bool { return a[j].Created.Before(a[i].Created) }
func (a test) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func NewestPost(posts []database.PostDisplay) {
	sort.Sort(test(posts))
}

// func MostComment(posts []database.PostDisplay) []database.PostDisplay {

// }

// func AllTime(posts []database.PostDisplay) []database.PostDisplay {

// }

// func YearAgo(posts []database.PostDisplay) []database.PostDisplay {

// }

// func MonthAgo(posts []database.PostDisplay) []database.PostDisplay {

// }

// func WeekAgo(posts []database.PostDisplay) []database.PostDisplay {

// }

// func DayAgo(posts []database.PostDisplay) []database.PostDisplay {

// }

// func HourAgo(posts []database.PostDisplay) []database.PostDisplay {

// }
