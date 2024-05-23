package inMemory

import (
	"errors"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

func (bd *InMemoryRepository) CreatePost(post *models.Post) (*models.Post, error) {
	bd.mutex.Lock()
	defer bd.mutex.Unlock()

	id := uuid.New()
	post.Id = id
	bd.PostRepository[id] = post

	return post, nil
}

func (bd *InMemoryRepository) DeletePostByID(postId uuid.UUID) error {
	bd.mutex.Lock()
	defer bd.mutex.Unlock()

	_, ok := bd.PostRepository[postId]
	if !ok {
		return errors.New("there is no such post")
	}

	for commentId := range bd.CommentRepository {
		if bd.CommentRepository[commentId].Post == postId {
			delete(bd.CommentRepository, commentId)
		}
	}

	delete(bd.PostRepository, postId)
	return nil
}
func (bd *InMemoryRepository) UpdatePost(post *models.Post) (*models.Post, error) {
	bd.mutex.Lock()
	defer bd.mutex.Unlock()

	_, ok := bd.PostRepository[post.Id]
	if !ok {
		return nil, errors.New("there is no such post")
	}

	bd.PostRepository[post.Id].Body = post.Body
	bd.PostRepository[post.Id].Title = post.Title
	bd.PostRepository[post.Id].Disabled = post.Disabled

	return bd.PostRepository[post.Id], nil
}

func (bd *InMemoryRepository) GetPostByID(postId uuid.UUID) (*models.Post, error) {

	bd.mutex.RLock()
	defer bd.mutex.RUnlock()

	post, ok := bd.PostRepository[postId]
	if !ok {
		return nil, errors.New("post not found")
	}

	return post, nil
}
