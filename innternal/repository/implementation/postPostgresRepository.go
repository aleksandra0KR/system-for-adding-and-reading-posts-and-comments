package implementation

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

type postPostgresRepository struct {
	db *pg.DB
}

func NewPostPostgresRepository(db *pg.DB) *postPostgresRepository {
	return &postPostgresRepository{db: db}
}

func (r *postPostgresRepository) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	_, err := r.db.WithContext(ctx).Model(post).Returning("*").Insert()

	if err != nil {
		return nil, err
	}
	fmt.Print(post)
	return post, nil
}

func (r *postPostgresRepository) DeletePostByID(ctx context.Context, id uuid.UUID) error {

	post := &models.Post{Id: id}
	_, err := r.db.WithContext(ctx).Model(post).Where("id = ?", post.Id).Delete()
	return err
}

func (r *postPostgresRepository) GetPostByID(ctx context.Context, id uuid.UUID) (*models.Post, error) {
	post := &models.Post{Id: id}

	err := r.db.WithContext(ctx).Model(&post).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postPostgresRepository) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	_, err := r.db.WithContext(ctx).Model(post).Where("id = ?", post.Id).Update()
	if err != nil {
		return nil, err
	}
	return post, nil

}
