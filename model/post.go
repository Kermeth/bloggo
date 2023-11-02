package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID      string `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created string    `json:"created"`
}

func New(title string, content string) Post {
	p := Post{uuid.New().String(), title, content, time.Now().Format("2006-01-02 15:04:05")}
	return p
}
