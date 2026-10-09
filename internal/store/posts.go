package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
	UserId    int64    `json:"userId"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	//inserting into the posts table
	query := `
	INSERT INTO posts (content, title, userId,  tags)
	VALUES ($1, $2, $3, $4) RETURNING id, createdAt, updatedAt
	`
	//accessing the store
	err := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserId, pq.Array(post.Tags)).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
