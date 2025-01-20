package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	if err = db.Ping(); err != nil {
		return nil
	}
	return db
}
