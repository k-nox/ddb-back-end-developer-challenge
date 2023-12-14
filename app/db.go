package app

import (
	"database/sql"
	"errors"
	"github.com/gocraft/dbr/v2"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

var UnexpectedDBError = errors.New("unexpected database error")

// applyMigrations uses takes in a directory path containing sql migration files and a *sql.DB instance, and applies any new migrations.
func applyMigrations(dir string, db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)

	if err != nil {
		log.Printf("error applying migrations: %s", err.Error())
		return UnexpectedDBError
	}

	log.Printf("Applied %d migrations\n", n)
	return nil
}

// openDB will open a new sqlite db at the provided dsn.
func openDB(dsn string) (*dbr.Connection, error) {
	conn, err := dbr.Open("sqlite3", dsn, nil)
	if err != nil {
		log.Printf("error opening db: %s", err.Error())
		return nil, UnexpectedDBError
	}
	return conn, nil
}

// CloseDB will close the app's sqlite database.
func (a *App) CloseDB() {
	err := a.db.Close()
	if err != nil {
		log.Fatalf("error closing db: %s", err.Error())
	}
}
