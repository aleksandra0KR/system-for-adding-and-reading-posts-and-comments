package inMemory

import (
	"github.com/google/uuid"
	"sync"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

type InMemoryRepository struct {
	PostRepository    map[uuid.UUID]*models.Post
	CommentRepository map[uuid.UUID]*models.Comment
	UserRepository    map[uuid.UUID]*models.User
	mutex             sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		PostRepository:    make(map[uuid.UUID]*models.Post),
		CommentRepository: make(map[uuid.UUID]*models.Comment),
		UserRepository:    make(map[uuid.UUID]*models.User),
	}
}
