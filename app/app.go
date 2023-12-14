// Package app provides the main database interactions for the API.
package app

import (
	"github.com/gocraft/dbr/v2"
)

// App is the main struct that maintains the connection to the database.
type App struct {
	db *dbr.Connection
}

// New takes in a dsn and a path ot the migrations and creates a new App instance with a sqlite db connection.
func New(dsn string, migrationsPath string) (*App, error) {
	// open db connection
	conn, err := openDB(dsn)
	if err != nil {
		return nil, err
	}

	// apply any migrations
	err = applyMigrations(migrationsPath, conn.DB)
	if err != nil {
		return nil, err
	}

	return &App{
		db: conn,
	}, nil
}
