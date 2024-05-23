package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"sync"
	"system-for-adding-and-reading-posts-and-comments/graph/model"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"

	"github.com/google/uuid"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post := &models.Post{
		Title:    input.Title,
		Body:     input.Body,
		UserId:   input.UserID,
		Disabled: input.Disabled,
	}
	post, err := r.Repository.CreatePost(post)
	if err != nil {
		return nil, err
	}
	postResult := &model.Post{
		ID:       post.Id,
		Title:    post.Title,
		Body:     post.Body,
		UserID:   post.UserId,
		Disabled: post.Disabled,
	}

	return postResult, err
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePost) (*model.Post, error) {
	post := &models.Post{
		Id:       input.ID,
		Title:    input.Title,
		Body:     input.Body,
		UserId:   input.UserID,
		Disabled: input.Disabled,
	}
	post, err := r.Repository.UpdatePost(post)
	if err != nil {
		return nil, err
	}
	result := &model.Post{
		ID:       post.Id,
		Title:    post.Title,
		Body:     post.Body,
		UserID:   post.UserId,
		Disabled: post.Disabled,
	}
	return result, err
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id uuid.UUID) (bool, error) {
	err := r.Repository.DeletePostByID(id)
	if err != nil {
		return false, err
	}
	return true, err
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	comment := &models.Comment{
		Body:   input.Body,
		UserId: input.UserID,
		Post:   input.PostID,
	}
	if input.ParentID == uuid.Nil {
		comment.Parent = uuid.Nil
	} else {
		comment.Parent = input.ParentID
	}
	comment, err := r.Repository.CreateComment(comment)
	if err != nil {
		return nil, err
	}
	commentResult := &model.Comment{
		ID:     comment.Id,
		Body:   comment.Body,
		UserID: comment.UserId,
		Parent: comment.Parent,
		Post:   comment.Post,
	}

	notify(comment.Post, commentResult)
	return commentResult, err
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, input *model.UpdateComment) (*model.Comment, error) {
	comment := &models.Comment{
		Id:   input.ID,
		Body: input.Body,
	}
	comment, err := r.Repository.UpdateComment(comment)
	if err != nil {
		return nil, err
	}
	result := &model.Comment{
		ID:     comment.Id,
		Body:   comment.Body,
		UserID: comment.UserId,
		Parent: comment.Parent,
		Post:   comment.Post,
	}
	return result, err
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id uuid.UUID) (bool, error) {
	err := r.Repository.DeleteCommentByID(id)
	if err != nil {
		return false, err
	}
	return true, err
}

// CreatUser is the resolver for the creatUser field.
func (r *mutationResolver) CreatUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &models.User{
		Name: input.Name,
	}
	user, err := r.Repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	commentResult := &model.User{
		ID:   user.Id,
		Name: user.Name,
	}

	return commentResult, err
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id uuid.UUID) (bool, error) {
	err := r.Repository.DeleteUserByID(id)
	if err != nil {
		return false, err
	}
	return true, err
}

// Post is the resolver for the Post field.
func (r *queryResolver) Post(ctx context.Context, postID uuid.UUID) (*model.Post, error) {
	post, err := r.Repository.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	comments, err := r.Repository.GetCommentsForPost(postID, 10, 0)
	if err != nil {
		return nil, err
	}
	resultPost := model.Post{
		ID:       post.Id,
		Title:    post.Title,
		Body:     post.Body,
		UserID:   post.UserId,
		Disabled: post.Disabled,
		Comments: comments,
	}
	return &resultPost, err

}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, limit *int, offset *int, postID uuid.UUID) ([]*model.Comment, error) {
	return r.Repository.GetCommentsForPost(postID, *limit, *offset)
}

var newCommentsChanel = make(map[uuid.UUID][]chan *model.Comment)
var mutex sync.Mutex

func notify(postID uuid.UUID, comment *model.Comment) {

	mutex.Lock()
	defer mutex.Unlock()
	if channels, found := newCommentsChanel[postID]; found {
		for _, ch := range channels {
			ch <- comment
		}
	}

}

// NewComment is the resolver for the newComment field.
func (r *subscriptionResolver) NewComment(ctx context.Context, postID uuid.UUID) (<-chan *model.Comment, error) {

	ch := make(chan *model.Comment, 1)
	mutex.Lock()
	newCommentsChanel[postID] = append(newCommentsChanel[postID], ch)
	mutex.Unlock()

	go func() {
		<-ctx.Done()
		mutex.Lock()
		defer mutex.Unlock()
		for i, c := range newCommentsChanel[postID] {
			if c == ch {
				newCommentsChanel[postID] = append(newCommentsChanel[postID][:i], newCommentsChanel[postID][i+1:]...)
				break
			}
		}
	}()

	return ch, nil

}

// Mutation returns MutationResolver inMemory.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver inMemory.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver inMemory.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error) {
	return r.Repository.GetPosts(*limit, *offset)
}
