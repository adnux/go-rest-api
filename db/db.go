package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SQLDB *sql.DB
var DB *gorm.DB

func InitDB() {
	var err error
	SQLDB, err := sql.Open("sqlite3", "api-gorm.db")
	SQLDB.SetMaxOpenConns(15)
	SQLDB.SetMaxIdleConns(3)
	if err != nil {
		panic("Could not connect to database.")
	}

	DB, err = gorm.Open(sqlite.New(sqlite.Config{
		Conn: SQLDB,
	}), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database.")
	}

}
