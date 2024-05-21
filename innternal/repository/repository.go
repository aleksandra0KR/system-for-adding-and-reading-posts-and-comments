package repository

import (
	"github.com/jmoiron/sqlx"
	"system-for-adding-and-reading-posts-and-comments/innternal/repository/implementation"
)

type Repository struct {
	UserRepository
	PostRepository
	CommentRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:    implementation.NewUserPostgresRepository(db),
		PostRepository:    implementation.NewPostPostgresRepository(db),
		CommentRepository: implementation.NewCommentPostgresRepository(db),
	}
}
