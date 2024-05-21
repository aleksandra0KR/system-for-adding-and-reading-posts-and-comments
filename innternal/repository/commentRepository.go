package repository

import (
	"context"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	DeleteCommentByID(ctx context.Context, id uuid.UUID) error
	GetCommentsForPost(ctx context.Context, id uuid.UUID, limit, offset int) ([]*models.Comment, error)
	UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	GetChildrenComments(ctx context.Context, id uuid.UUID) ([]*models.Comment, error)
}
