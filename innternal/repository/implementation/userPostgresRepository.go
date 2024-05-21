package implementation

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"system-for-adding-and-reading-posts-and-comments/innternal/model"
)

type userPostgresRepository struct {
	db *sqlx.DB
}

func NewUserPostgresRepository(db *sqlx.DB) *userPostgresRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) CreateUser(user *model.User) error {
	var id uuid.UUID
	query := `INSERT INTO "users" ("Name") VALUES ($1) RETURNING userId`

	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return err
	}
	user.UserId = id

	return nil
}

func (r *userPostgresRepository) DeleteUserByID(id uuid.UUID) error {

	query := `DELETE FROM "users" where "id"= $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
