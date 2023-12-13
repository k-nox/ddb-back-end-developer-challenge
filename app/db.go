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

func openDB(dsn string) (*dbr.Connection, error) {
	conn, err := dbr.Open("sqlite3", dsn, nil)
	if err != nil {
		log.Printf("error opening db: %s", err.Error())
		return nil, UnexpectedDBError
	}
	return conn, nil
}

func (a *App) CloseDB() {
	err := a.db.Close()
	if err != nil {
		log.Fatalf("error closing db: %s", err.Error())
	}
}
