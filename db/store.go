package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func OpenDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Printf("%s", err)
		return nil
	}

	query_statement := `
		PRAGMA JOURNAL_MODE = WAL;
		PRAGMA FOREIGN_KEYS = ON;
		PRAGMA BUSY_TIMEOUT = 500;

	`

	_, err = db.Exec(query_statement)

	if err != nil {
		log.Printf("%s", err)
	}
	query_statement = `
		CREATE TABLE IF NOT EXISTS profiles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			master_password TEXT
		);
	`
	_, err = db.Exec(query_statement)

	if err != nil {
		log.Printf("%s", err)
	}

	query_statement = `
		CREATE TABLE IF NOT EXISTS password_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user TEXT,
			app TEXT,
			password TEXT,
			note TEXT
		);
	`
	_, err = db.Exec(query_statement)

	if err != nil {
		log.Printf("%s", err)
	}

	return db

}
