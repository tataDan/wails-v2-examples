package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connect(username string, password string) error {
	connStr := fmt.Sprintf("%s:%s@tcp(localhost:3306)/nation",
		username, password)
	var err error

	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func getDb() *sql.DB {
	return db
}
