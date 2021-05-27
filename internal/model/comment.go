package model

import (
	"time"
)

// Comment defines a comment object
type Comment struct {
	ID       uint      `json:"id"`
	Body     string    `json:"body"`
	Author   string    `json:"author"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

// GetCommentsQuery returns query for getting all comments from db
func (c *Comment) GetCommentsQuery() string {
	return `SELECT * FROM comments;`
}

// GetCommentQuery returns query for getting a single comment from db
func (c *Comment) GetCommentQuery() string {
	return `SELECT * FROM comments WHERE id=$1;`
}
