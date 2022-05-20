package store

import "github.com/hono3bono3/go-plactice/server/handler"

var comments = make([]handler.Comment, 0, 100)

func GetComments() []handler.Comment {
	return comments
}
