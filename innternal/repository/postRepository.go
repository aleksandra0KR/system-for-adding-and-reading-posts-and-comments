package repository

import (
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/model"
)

type PostRepository interface {
	CreatePost(post *model.Post) error
	DeletePostByID(id uuid.UUID) error
	GetPostByID(id uuid.UUID) (*model.Post, error)
	UpdatePost(post *model.Post) error
}
