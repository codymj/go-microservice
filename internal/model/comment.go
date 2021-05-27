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
	return `SELECT * FROM comments 
			WHERE id=$1;`
}

// CreateCommentQuery returns query for inserting a new comment into db
func (c *Comment) CreateCommentQuery() string {
	return `INSERT INTO comments (body, author, created, modified) 
			VALUES ($1, $2, $3, $4)`
}

// UpdateCommentQuery returns query for updating a row by ID in db
func (c *Comment) UpdateCommentQuery() string {
	return `UPDATE comments 
			SET body=$2, author=$3, modified=$4 
			WHERE id=$1`
}

// DeleteCommentQuery returns query for deleting a row by ID in db
func (c *Comment) DeleteCommentQuery() string {
	return `DELETE FROM comments 
			WHERE id=$1`
}
