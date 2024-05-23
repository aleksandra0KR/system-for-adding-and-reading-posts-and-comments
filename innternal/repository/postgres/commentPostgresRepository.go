package postgres

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/graph/model"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

func (r *Repository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	if len(comment.Body) > 2000 {
		return nil, errors.New("body of the comment is too long")
	}
	row := r.db.QueryRow(`SELECT "disabled" FROM "posts" WHERE "id" = $1`, comment.Post)

	var allowedComments bool
	if err := row.Scan(&allowedComments); err != nil {
		return nil, err
	}
	if !allowedComments {
		return nil, errors.New("comments are not allowed")
	}

	var id uuid.UUID
	if comment.Parent == uuid.Nil {
		query := `INSERT INTO "comments" ("body", "user_id", "post") VALUES ($1, $2, $3) RETURNING id`
		row := r.db.QueryRow(query, comment.Body, comment.UserId, comment.Post)
		if err := row.Scan(&id); err != nil {
			return nil, err
		}
	} else {
		query := `INSERT INTO "comments" ("body", "user_id", "parent", "post") VALUES ($1, $2, $3, $4) RETURNING id`
		row := r.db.QueryRow(query, comment.Body, comment.UserId, comment.Parent, comment.Post)
		if err := row.Scan(&id); err != nil {
			return nil, err
		}
	}

	comment.Id = id

	return comment, nil
}

func (r *Repository) DeleteCommentByID(id uuid.UUID) error {

	query := `DELETE FROM "comments" where "id"= $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (r *Repository) UpdateComment(comment *models.Comment) (*models.Comment, error) {
	if len(comment.Body) > 2000 {
		return nil, errors.New("body of the comment is too long")
	}

	query, err := r.db.Prepare(`UPDATE "comments" SET "body"=$1 WHERE "id"=$2`)
	if err != nil {
		return nil, err
	}
	_, err = query.Exec(comment.Body, comment.Id)

	row := r.db.QueryRow(`SELECT "body", "user_id", "parent", "post" FROM "comments" WHERE "id" = $1`, comment.Id)

	if err := row.Scan(&comment.Body, &comment.UserId, &comment.Parent, &comment.Post); err != nil {
		return nil, err
	}

	return comment, err
}

func (r *Repository) GetCommentsForPost(id uuid.UUID, limit, offset int) ([]*model.Comment, error) {

	rows, err := r.db.Query(`SELECT "id", "body", "user_id", "parent", "post" FROM "comments" WHERE post = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentMap := make(map[uuid.UUID]*model.Comment)

	for rows.Next() {
		var c model.Comment
		if err := rows.Scan(&c.ID, &c.Body, &c.UserID, &c.Parent, &c.Post); err != nil {
			return nil, err
		}
		commentMap[c.ID] = &c
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for _, comment := range commentMap {
		if *comment.Parent != uuid.Nil {
			parentComment, ok := commentMap[*comment.Parent]
			if ok {
				parentComment.Children = append(parentComment.Children, comment)
			}
		}
	}

	var comments []*model.Comment
	if limit == 0 {
		return comments, nil
	}

	start := offset
	end := limit + offset

	if end > len(commentMap) {
		end = len(commentMap)
	}

	for _, comment := range commentMap {
		fmt.Println(comment)
		if len(comments) >= end {
			return comments, nil
		}
		comments = append(comments, comment)

	}

	return comments[start:end], nil

}
