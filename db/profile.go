package db

import (
	"context"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EditProfile(database *sql.DB, profile Profile) error {
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

func DeleteProfile(database *sql.DB, profile Profile) error {

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

func AddProfile(database *sql.DB, profile Profile) error {
	statement := `INSERT INTO profiles (name , master_password) VALUES(?,?);`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(profile.MasterPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = database.Exec(statement, profile.Name, hashedPassword)
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