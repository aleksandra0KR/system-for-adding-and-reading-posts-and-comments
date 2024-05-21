package implementation

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

type userPostgresRepository struct {
	db *pg.DB
}

func NewUserPostgresRepository(db *pg.DB) *userPostgresRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := r.db.WithContext(ctx).Model(user).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userPostgresRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {

	user := &models.User{Id: id}
	_, err := r.db.WithContext(ctx).Model(user).Where("id = ?", user.Id).Delete()
	return err
}
