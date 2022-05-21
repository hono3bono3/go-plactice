package store

import (
	"github.com/hono3bono3/go-plactice/models"
)

var comments = make([]models.Comment, 0, 100)

func GetComments() []models.Comment {
	return comments
}

func AddComment(c models.Comment) {
	comments = append(comments, c)
}
