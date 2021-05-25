package model

import (
	"time"
)

// Comment defines a comment object
type Comment struct {
	Id			uint		`json:"id"`
	Body		string		`json:"body"`
	Author		string		`json:"author"`
	Created		time.Time	`json:"created"`
	Modified	time.Time	`json:"modified"`
}
