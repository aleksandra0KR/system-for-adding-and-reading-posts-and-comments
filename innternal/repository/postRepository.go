package repository

import (
	"context"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	DeletePostByID(ctx context.Context, id uuid.UUID) error
	GetPostByID(ctx context.Context, id uuid.UUID) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
}
