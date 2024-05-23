package inMemory

import (
	"errors"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

func (bd *InMemoryRepository) CreateUser(user *models.User) (*models.User, error) {
	bd.mutex.Lock()
	defer bd.mutex.Unlock()

	id := uuid.New()
	user.Id = id
	bd.UserRepository[id] = user

	return user, nil
}

func (bd *InMemoryRepository) DeleteUserByID(userId uuid.UUID) error {
	bd.mutex.Lock()
	defer bd.mutex.Unlock()

	_, ok := bd.UserRepository[userId]
	if !ok {
		return errors.New("there is no such user")
	}

	for commentId := range bd.CommentRepository {
		if bd.CommentRepository[commentId].UserId == userId {
			delete(bd.CommentRepository, commentId)
		}
	}

	for postId := range bd.PostRepository {
		if bd.PostRepository[postId].UserId == userId {
			delete(bd.PostRepository, postId)
		}
	}

	delete(bd.UserRepository, userId)
	return nil
}
