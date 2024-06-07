package api

import (
	"ThreadCore/database"
	"sort"
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
