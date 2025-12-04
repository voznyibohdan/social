package storage

import "database/sql"

type Storage struct {
	Users UsersStorage
	Posts PostsStorage
}

func NewPostgresStorage(db *sql.DB) *Storage {
	return &Storage{
		Users: NewPostgresUsersStorage(db),
		Posts: NewPostgresPostsStorage(db),
	}
}
