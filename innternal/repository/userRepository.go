package repository

import (
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	DeleteUserByID(id uuid.UUID) error
}
