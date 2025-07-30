package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=yourpassword dbname=assetdb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS assets (
		id SERIAL PRIMARY KEY,
		name TEXT,
		category TEXT,
		status TEXT
	)`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
