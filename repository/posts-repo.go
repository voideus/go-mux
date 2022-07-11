package repository

import (
	"github.com/voideus/golang-mux-rest/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
