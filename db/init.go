package db

import (
	"context"
	"database/sql"
	_ "embed"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var DDL string

var DB *sql.DB
var DBQueries *Queries
var CTX = context.Background()

func InitDB() error {

	databaseName := os.Getenv("DATABASE_NAME")

	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(CTX, DDL); err != nil {
		return err
	}

	DB = db
	DBQueries = New(db)

	return nil
}
