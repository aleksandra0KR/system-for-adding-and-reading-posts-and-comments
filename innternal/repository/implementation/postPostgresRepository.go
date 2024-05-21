package implementation

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"system-for-adding-and-reading-posts-and-comments/innternal/model"
)

type postPostgresRepository struct {
	db *sqlx.DB
}

func NewPostPostgresRepository(db *sqlx.DB) *postPostgresRepository {
	return &postPostgresRepository{db: db}
}

func (r *postPostgresRepository) CreatePost(post *model.Post) error {
	var id uuid.UUID
	query := `INSERT INTO "posts" ("title", "body", "userId","disabledComments") VALUES ($1, $2, $3, $4) RETURNING postId`
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return err
	}
	post.PostId = id

	return nil
}

func (r *postPostgresRepository) DeletePostByID(id uuid.UUID) error {

	query := `DELETE FROM "posts" where "postId"= $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *postPostgresRepository) GetPostByID(id uuid.UUID) (*model.Post, error) {
	query := `SELECT * FROM "posts" WHERE "postId" = $1`

	var post model.Post
	err := r.db.Get(&post, query, id)

	return &post, err
}

func (r *postPostgresRepository) UpdatePost(post *model.Post) error {
	query := ` UPDATE "posts" SET "title" = $2, "body" = $3, "disabledComments" = $4  WHERE "postId" = $1`
	_, err := r.db.Exec(query, post.PostId, post.Title, post.Body, post.DisabledComments)
	return err
}
