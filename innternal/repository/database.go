package repository

import (
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/graph/model"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

type Database interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	DeleteCommentByID(id uuid.UUID) error
	GetCommentsForPost(id uuid.UUID, limit, offset int) ([]*model.Comment, error)
	UpdateComment(comment *models.Comment) (*models.Comment, error)
	CreatePost(post *models.Post) (*models.Post, error)
	DeletePostByID(id uuid.UUID) error
	GetPostByID(id uuid.UUID) (*models.Post, error)
	GetPosts(limit, offset int) ([]*model.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
	CreateUser(user *models.User) (*models.User, error)
	DeleteUserByID(id uuid.UUID) error
}
