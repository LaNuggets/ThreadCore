package api

import (
	"ThreadCore/database"
	"math"
	"sort"
	"strconv"
	"time"
)

type post []database.PostInfo

func (a post) Len() int {
	return len(a)
}
func (a post) Less(i, j int) bool { return a[j].Created.Before(a[i].Created) }
func (a post) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func NewestPost(posts []database.PostInfo) {
	sort.Sort(post(posts))
}

func GetFormatedDuration(duration time.Duration) string {
	if math.Trunc(duration.Hours()/24/365) != 0 {
		return strconv.Itoa(int(math.Trunc(duration.Hours()/24/365))) + "yr. " + strconv.Itoa(int(math.Trunc(duration.Hours()/24/30))-(int(math.Trunc(duration.Hours()/24/365))*12)) + "mo. ago"
	} else if math.Trunc(duration.Hours()/24/30) != 0 {
		return strconv.Itoa(int(math.Trunc(duration.Hours()/24/30))) + "mo. " + strconv.Itoa(int(math.Trunc(math.Trunc(duration.Hours()/24/7)-(math.Trunc(duration.Hours()/24/30))*(30/7)))) + "w. ago"
	} else if math.Trunc(duration.Hours()/24/7) != 0 {
		return strconv.Itoa(int(math.Trunc(duration.Hours()/24/7))) + "w. " + strconv.Itoa(int(math.Trunc(duration.Hours()/24))-(int(math.Trunc(duration.Hours()/24/7))*7)) + "d. ago"
	} else if math.Trunc(duration.Hours()/24) != 0 {
		return strconv.Itoa(int(math.Trunc(duration.Hours()/24))) + "d. " + strconv.Itoa(int(math.Trunc(duration.Hours()))-(int(math.Trunc(duration.Hours()/24))*24)) + "h. ago"
	} else if math.Trunc(duration.Hours()) != 0 {
		return strconv.Itoa(int(math.Trunc(duration.Hours()))) + "h. " + strconv.Itoa(int(math.Trunc(duration.Minutes()))-(int(math.Trunc(duration.Hours()))*60)) + "m. ago"
	} else if math.Trunc(duration.Minutes()) != 0 {
		return strconv.Itoa(int(math.Trunc(duration.Minutes()))) + "m. " + strconv.Itoa(int(math.Trunc(duration.Seconds()))-(int(math.Trunc(duration.Minutes()))*60)) + "s. ago"
	} else {
		return strconv.Itoa(int(math.Trunc(duration.Seconds()))) + "s. ago"
	}
}
