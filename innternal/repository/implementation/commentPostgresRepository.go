package implementation

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

type commentPostgresRepository struct {
	db *pg.DB
}

func NewCommentPostgresRepository(db *pg.DB) *commentPostgresRepository {
	return &commentPostgresRepository{db: db}
}

func (r *commentPostgresRepository) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	_, err := r.db.WithContext(ctx).Model(comment).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentPostgresRepository) DeleteCommentByID(ctx context.Context, id uuid.UUID) error {

	comment := &models.Comment{Id: id}
	_, err := r.db.WithContext(ctx).Model(comment).Where("id = ?", comment.Id).Delete()
	return err

}

func (r *commentPostgresRepository) UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	_, err := r.db.WithContext(ctx).Model(comment).Where("id = ?", comment.Id).Update()
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentPostgresRepository) GetCommentsForPost(ctx context.Context, id uuid.UUID, limit, offset int) ([]*models.Comment, error) {
	var comments []*models.Comment

	query := r.db.WithContext(ctx).Model(&comments).Order("id")
	query.Limit(limit)
	query.Offset(offset)

	err := query.Select()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentPostgresRepository) GetChildrenComments(ctx context.Context, id uuid.UUID) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := r.db.WithContext(ctx).Model(&comments).Where("parent = ?", id).Order("id").Select()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
