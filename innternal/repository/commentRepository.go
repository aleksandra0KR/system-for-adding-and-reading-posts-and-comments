package repository

import (
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/model"
)

type CommentRepository interface {
	CreateComment(comment *model.Comment) error
	DeleteCommentByID(id uuid.UUID) error
	GetCommentsForPost(id uuid.UUID) ([]*model.Comment, error)
	UpdateComment(comment *model.Comment) error
	GetChildrenComments(id uuid.UUID) ([]*model.Comment, error)
}
