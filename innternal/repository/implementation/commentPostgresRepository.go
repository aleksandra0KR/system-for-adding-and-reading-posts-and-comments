package implementation

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"system-for-adding-and-reading-posts-and-comments/innternal/model"
)

type commentPostgresRepository struct {
	db *sqlx.DB
}

func NewCommentPostgresRepository(db *sqlx.DB) *commentPostgresRepository {
	return &commentPostgresRepository{db: db}
}

func (r *commentPostgresRepository) CreateComment(comment *model.Comment) error {
	var id uuid.UUID
	query := `INSERT INTO "comments" ("body", "userId","parent", "postId") VALUES ($1, $2, $3, $4) RETURNING commentId`
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return err
	}
	comment.CommentId = id

	return nil
}

func (r *commentPostgresRepository) DeleteCommentByID(id uuid.UUID) error {

	query := `DELETE FROM "comments" where "commentId"= $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *commentPostgresRepository) UpdateComment(comment *model.Comment) error {
	query := ` UPDATE "comments" SET "body" = $2 WHERE "commentId" = $1`
	_, err := r.db.Exec(query, comment.CommentId, comment.Body)
	return err
}

func (r *commentPostgresRepository) GetCommentsForPost(id uuid.UUID) ([]*model.Comment, error) {
	query := ` SELECT * FROM  "comments" WHERE "postId" = $1`

	var comments []*model.Comment
	rows, err := r.db.Query(query, id)
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var c model.Comment
		err = rows.Scan(&c.CommentId, &c.Body, &c.UserId, &c.Parent, &c.Post)

		if err != nil {
			return comments, err
		}

		comments = append(comments, &c)

	}
	return comments, nil
}

func (r *commentPostgresRepository) GetChildrenComments(id uuid.UUID) ([]*model.Comment, error) {
	query := ` SELECT * FROM  "comments" WHERE "parent" = $1`

	var comments []*model.Comment
	rows, err := r.db.Query(query, id)
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var c model.Comment
		err = rows.Scan(&c.CommentId, &c.Body, &c.UserId, &c.Parent, &c.Post)

		if err != nil {
			return comments, err
		}

		comments = append(comments, &c)

	}
	return comments, nil
}
