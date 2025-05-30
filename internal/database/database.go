package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(user, pass, host string, port int, name string) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		pass,
		host,
		port,
		name,
	)

	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return db, nil
}
