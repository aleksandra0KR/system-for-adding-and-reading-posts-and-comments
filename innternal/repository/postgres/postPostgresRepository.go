package postgres

import (
	"github.com/google/uuid"
	"system-for-adding-and-reading-posts-and-comments/graph/model"
	"system-for-adding-and-reading-posts-and-comments/innternal/models"
)

func (r *Repository) CreatePost(post *models.Post) (*models.Post, error) {

	var id uuid.UUID
	query := `INSERT INTO "posts" ("title", "body", "user_id", "disabled") VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.db.QueryRow(query, post.Title, post.Body, post.UserId, post.Disabled)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	post.Id = id

	return post, nil
}

func (r *Repository) DeletePostByID(id uuid.UUID) error {

	query := `DELETE FROM "posts" where "id"= $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (r *Repository) GetPostByID(id uuid.UUID) (*models.Post, error) {
	row := r.db.QueryRow(`SELECT "id", "title", "body", "user_id", "disabled" FROM "posts" WHERE "id" = $1`, id)

	var post models.Post
	if err := row.Scan(&post.Id, &post.Title, &post.Body, &post.UserId, &post.Disabled); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *Repository) UpdatePost(post *models.Post) (*models.Post, error) {
	query, err := r.db.Prepare(`UPDATE "posts" SET "title"=$1 WHERE "id"=$2`)
	if err != nil {
		return nil, err
	}
	_, err = query.Exec(post.Title, post.Id)
	if err != nil {
		return nil, err
	}

	query2, err := r.db.Prepare(`UPDATE "posts" SET "body"=$1  WHERE "id"=$2`)
	if err != nil {
		return nil, err
	}
	_, err = query2.Exec(post.Body, post.Id)

	query3, err := r.db.Prepare(`UPDATE "posts" SET  "disabled"=$1  WHERE "id"=$2`)
	if err != nil {
		return nil, err
	}
	_, err = query3.Exec(post.Disabled, post.Id)

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(`SELECT "title", "body", "user_id", "disabled" FROM "posts" WHERE "id" = $1`, post.Id)

	if err := row.Scan(&post.Title, &post.Body, &post.UserId, &post.Disabled); err != nil {
		return nil, err
	}

	return post, err
}

func (r *Repository) GetPosts(limit, offset int) ([]*model.Post, error) {

	rows, err := r.db.Query(`SELECT "id", "title", "body", "user_id" FROM "posts"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	postMap := make(map[uuid.UUID]*model.Post)

	for rows.Next() {
		var p model.Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.UserID); err != nil {
			return nil, err
		}
		postMap[p.ID] = &p
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var posts []*model.Post
	if limit == 0 {
		return posts, nil
	}

	start := offset
	end := limit + offset

	if end > len(postMap) {
		end = len(postMap)
	}

	for _, post := range postMap {
		if len(posts) >= end {
			return posts, nil
		}
		posts = append(posts, post)

	}

	return posts[start:end], nil
}
