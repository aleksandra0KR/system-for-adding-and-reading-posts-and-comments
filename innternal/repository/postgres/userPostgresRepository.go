package postgres

import (
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

func (r *Repository) CreateUser(user *models.User) (*models.User, error) {
	var id uuid.UUID
	query := `INSERT INTO "users" ("name") VALUES ($1) RETURNING id`

	row := r.db.QueryRow(query, user.Name)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	user.Id = id

	return user, nil
}

func (r *Repository) DeleteUserByID(id uuid.UUID) error {

	query := `DELETE FROM "users" where "id"= $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
