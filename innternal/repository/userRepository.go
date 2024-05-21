package repository

import (
	"context"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUserByID(ctx context.Context, id uuid.UUID) error
}
