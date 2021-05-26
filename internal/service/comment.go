package service

import (
	"codymj/go-microservice/internal/model"
	"database/sql"
	"errors"
	"log"
)

// Service provides database access for a service
type Service struct {
	DB *sql.DB
}

// New returns a pointer to base service with db access
func New() *Service {
	return &Service{}
}

// CommentService defines the comment service contract
type CommentService interface {
	GetComments() ([]model.Comment, error)
	GetComment(id uint) (model.Comment, error)
	CreateComment(comment model.Comment) (model.Comment, error)
	UpdateComment(id uint, comment model.Comment) (model.Comment, error)
	DeleteComment(id uint) error
}

// GetComment returns a single comment
func (s *Service) GetComment(id uint) (model.Comment, error) {
	var comment model.Comment

	query := comment.GetCommentQuery()
	row := s.DB.QueryRow(query, id)

	err := row.Scan(
		&comment.ID, &comment.Body, &comment.Author,
		&comment.Created, &comment.Modified,
	)
	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned")
		return comment, errors.New("no rows were returned")
	case nil:
		return comment, nil
	default:
		panic(err)
	}
}
