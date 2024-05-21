package repository

import (
	"github.com/go-pg/pg/v10"
	"system-for-adding-and-reading-posts-and-comments/innternal/repository/implementation"
)

type Repository struct {
	UserRepository
	PostRepository
	CommentRepository
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		UserRepository:    implementation.NewUserPostgresRepository(db),
		PostRepository:    implementation.NewPostPostgresRepository(db),
		CommentRepository: implementation.NewCommentPostgresRepository(db),
	}
}
