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
	GetOneByID(context.Context, int64) (*Post, error)
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

func (s *PostgresPostsStorage) GetOneByID(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT id, user_id, content, title, created_at, updated_at, tags 
		FROM posts 
		WHERE id = $1`

	var post Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Content,
		&post.Title,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
	)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
