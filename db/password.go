package db

import (
	"context"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EdtPasswordItem(database *sql.DB, profile PasswordItem) error {
	statement := `UPDATE profiles SET name = ?, master_password = ? WHERE id = ?;`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(profile.MasterPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = database.Exec(statement, profile.Name, hashedPassword, profile.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetPassword(database *sql.DB, id int) (PasswordItem, error) {
	profile := PasswordItem{}
	statement := "SELECT * FROM password_items WHERE ID = ?"
	result := database.QueryRow(statement, id)

	err := result.Scan(&profile.Id, &profile.User, &profile.App, &profile.Password, &profile.Note)

	if err != nil {
		return profile, err
	}
	return profile, nil
}

func EditPasswordItem(database *sql.DB, password PasswordItem) error {
	statement := `UPDATE password_items SET app = ?, password = ?, note = ? WHERE id = ?;`

	return nil
}

func DeeteProfile(database *sql.DB, profile Profile) error {

	tx, err := database.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM password_items WHERE user = ?;", profile.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM profiles WHERE id = ?;", profile.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func AddPasswordItem(database *sql.DB, password PasswordItem) error {
	statement := `INSERT INTO password_items (user, app, password, note) VALUES(?,?,?,?) `

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = database.Exec(statement, password.User, password.App, hashedPassword, password.Note)
	if err != nil {
		return err
	}

	return nil
}

func GeProfiles(database *sql.DB) ([]Profile, error) {
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
