package post

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Created string `json:"created"`
}

func New(title string, content string) *Post {
	p := &Post{
		uuid.NewString(),
		title,
		content,
		time.Now().Format(time.RFC3339),
	}
	return p
}
