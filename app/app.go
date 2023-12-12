package app

import (
	"github.com/gocraft/dbr/v2"
)

type App struct {
	db *dbr.Connection
}

func New() (*App, error) {
	// open db connection
	conn, err := openDB("db/data.db")
	if err != nil {
		return nil, err
	}

	// apply any migrations
	err = applyMigrations("db/migrations", conn.DB)
	if err != nil {
		return nil, err
	}

	return &App{
		db: conn,
	}, nil
}
