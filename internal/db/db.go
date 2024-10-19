package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() error {
	var err error
	dsn := "user:password@tcp(localhost:3306)/userdb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
