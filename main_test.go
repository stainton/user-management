package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(t *testing.T) {
	// Test cases for main.go go here
	db, err := sql.Open("mysql", "root:961110@tcp(localhost:3306)/testdb")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	// query all databses in the db connection
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	var dbName string
	for rows.Next() {
		err := rows.Scan(&dbName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(dbName)
	}
}
