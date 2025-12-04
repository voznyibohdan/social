package db

import (
	"context"
	"database/sql"
	"time"
)

func OpenDB(dsn string, maxIdleTime time.Duration, maxOpenConns, maxIdleConns int) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxIdleConns)
	db.SetMaxIdleConns(maxOpenConns)
	db.SetConnMaxIdleTime(maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}
