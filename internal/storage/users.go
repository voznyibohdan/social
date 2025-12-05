package storage

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UsersStorage interface {
	Create(context.Context, *User) error
}

type PostgresUsersStorage struct {
	db *sql.DB
}

func NewPostgresUsersStorage(db *sql.DB) *PostgresUsersStorage {
	return &PostgresUsersStorage{db: db}
}

func (s *PostgresUsersStorage) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (username, email, password) 
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at`

	args := []any{user.Username, user.Email, user.Password}

	err := s.db.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
