package db

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var DDL string

var DB *sql.DB
var DBQueries *Queries
var CTX = context.Background()

func InitDB() error {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(CTX, DDL); err != nil {
		return err
	}

	DBQueries = New(db)

	return nil
}
