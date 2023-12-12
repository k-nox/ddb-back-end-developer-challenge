package app

import (
	"database/sql"
	"fmt"
	"github.com/gocraft/dbr/v2"
	_ "github.com/mattn/go-sqlite3"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func applyMigrations(dir string, db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)

	if err != nil {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	log.Printf("Applied %d migrations!\n", n)
	return nil
}

func openDB(location string) (*dbr.Connection, error) {
	conn, err := dbr.Open("sqlite3", location, nil)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}
	return conn, nil
}

func (a *App) CloseDB() {
	err := a.db.Close()
	if err != nil {
		log.Fatalf("error closing db: %s", err.Error())
	}
}
