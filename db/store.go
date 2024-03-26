package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
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

func AddProfile(database *sql.DB, profile Profile) error {
	statement := `INSERT INTO profiles (name , master_password) VALUES(?,?);`
	_, err := database.Exec(statement, profile.Name, profile.MasterPassword)
	if err != nil {
		return err
	}
	return nil
}

func GetProfiles(database *sql.DB) ([]Profile, error) {
	var profiles []Profile
	statement := "SELECT * FROM profiles"
	result, err := database.Query(statement)
	if err != nil {
		return profiles, err
	}
	defer result.Close()

	for result.Next() {
		var profile Profile
		err = result.Scan(&profile.Id, &profile.Name, &profile.MasterPassword)
		log.Printf("%v", profile)
		if err != nil {
			continue
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}
