package storage

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64    `json:"id"`
	UserID    int64    `json:"user_id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Tags      []string `json:"tags"`
}

type PostsStorage interface {
	Create(context.Context, *Post) error
}

type PostgresPostsStorage struct {
	db *sql.DB
}

func NewPostgresPostsStorage(db *sql.DB) *PostgresPostsStorage {
	return &PostgresPostsStorage{db: db}
}

func (s *PostgresPostsStorage) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (title, content, user_id, tags) 
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at`

	args := []any{post.Title, post.Content, post.UserID, pq.Array(post.Tags)}

	err := s.db.QueryRowContext(ctx, query, args...).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
