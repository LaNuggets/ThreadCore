package api

import "ThreadCore/database"

func DisplayPosts(post []database.Post) []string {
	var commentContent []string
	for i := 0; i < len(post); i++ {
		commentContent = append(commentContent, post[i].Content)
	}
	return commentContent
}
