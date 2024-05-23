package inMemory

import (
	"errors"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/graph/model"
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

func (bd *InMemoryRepository) GetPosts(limit int, offset int) ([]*model.Post, error) {
	bd.mutex.RLock()
	defer bd.mutex.RUnlock()

	var posts []*model.Post
	for _, post := range bd.PostRepository {

		posts = append(posts, &model.Post{
			Title:  post.Title,
			Body:   post.Body,
			UserID: post.UserId,
		})

	}

	if offset > len(posts) {
		return []*model.Post{}, errors.New("offset out of range")
	}
	limit = limit + offset
	if limit > len(posts) {
		limit = len(posts)
	}
	return posts[offset:limit], nil
}
