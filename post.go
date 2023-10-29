package main

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created string    `json:"created"`
}

func createPost(title string, content string) Post {
	p := Post{uuid.New(), title, content, time.Now().Format("2006-01-02 15:04:05")}
	return p
}
