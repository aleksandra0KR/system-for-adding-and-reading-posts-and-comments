package inMemory

import (
	"errors"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/graph/model"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

func (bd *InMemoryRepository) GetCommentsForPost(postId uuid.UUID, limit int, offset int) ([]*model.Comment, error) {
	bd.mutex.RLock()
	defer bd.mutex.RUnlock()

	var comments []*model.Comment
	for _, comment := range bd.CommentRepository {
		if comment.Post == postId {
			comments = append(comments, &model.Comment{
				Body:   comment.Body,
				UserID: comment.UserId,
				Post:   comment.Post,
				Parent: &comment.Parent,
			})
		}
	}

	if offset > len(comments) {
		return []*model.Comment{}, errors.New("offset out of range")
	}
	limit = limit + offset
	if limit > len(comments) {
		limit = len(comments)
	}
	return comments[offset:limit], nil
}

func (bd *InMemoryRepository) DeleteCommentByID(commentId uuid.UUID) error {
	bd.mutex.Lock()
	defer bd.mutex.Unlock()

	_, ok := bd.CommentRepository[commentId]
	if !ok {
		return errors.New("there is no such comment")
	}

	delete(bd.CommentRepository, commentId)
	return nil
}

func (bd *InMemoryRepository) UpdateComment(comment *models.Comment) (*models.Comment, error) {
	bd.mutex.Lock()
	defer bd.mutex.Unlock()

	_, ok := bd.PostRepository[comment.Id]
	if !ok {
		return nil, errors.New("there is no such comment")
	}
	if len(comment.Body) > 2000 {
		return nil, errors.New("body of the comment is too long")
	}

	bd.CommentRepository[comment.Id].Body = comment.Body
	return bd.CommentRepository[comment.Id], nil
}

func (bd *InMemoryRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	bd.mutex.Lock()
	defer bd.mutex.Unlock()

	if len(comment.Body) > 2000 {
		return nil, errors.New("body of the comment is too long")
	}

	post, ok := bd.PostRepository[comment.Post]
	if !ok {
		return nil, errors.New("there is no such post")
	}
	if !post.Disabled {
		return nil, errors.New("comments are not allowed")
	}

	id := uuid.New()
	comment.Id = id
	if comment.Parent != uuid.Nil {
		parentComment, ok := bd.CommentRepository[comment.Parent]
		if !ok {
			return nil, errors.New("parent comment not found")
		}
		parentComment.Children = append(parentComment.Children, comment)
	}

	post.Comments = append(post.Comments, comment)
	bd.CommentRepository[id] = comment
	return comment, nil
}
