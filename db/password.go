package db

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func SearchPasswordItems(database *sql.DB, user, query string) []PasswordItem {
	var passwordItems []PasswordItem
	statement := "SELECT * FROM password_items WHERE app LIKE ? AND user = ?"
	result, err := database.Query(statement, "%"+query+"%", user)
	defer result.Close()
	if err != nil {
		return passwordItems
	}

	for result.Next() {
		var password PasswordItem
		err = result.Scan(&password.Id, &password.User, &password.App, &password.Password, &password.Note)
		if err != nil {
			continue
		}
		passwordItems = append(passwordItems, password)
	}

	return passwordItems
}

func GetPassword(database *sql.DB, id int) (PasswordItem, error) {
	password := PasswordItem{}
	statement := "SELECT * FROM password_items WHERE ID = ?"
	result := database.QueryRow(statement, id)

	err := result.Scan(&password.Id, &password.User, &password.App, &password.Password, &password.Note)

	if err != nil {
		return password, err
	}
	return password, nil
}

func EditPasswordItem(database *sql.DB, password PasswordItem) error {
	statement := `UPDATE password_items SET app = ?, password = ?, note = ? WHERE id = ?;`
	previousPassword, err := GetPassword(database, int(password.Id))
	if err != nil {
		return err
	}
	previousPassword.CompareAndEdit(password)
	_, err = database.Exec(statement, previousPassword.App, previousPassword.Password, previousPassword.Note, previousPassword.Id)
	if err != nil {
		return err
	}

	return nil
}

func DeletePasswordItem(database *sql.DB, password PasswordItem) error {
	statement := `DELETE FROM password_items WHERE id = ?;`
	_, err := database.Exec(statement, password.Id)
	if err != nil {
		return err
	}
	return nil

}

func AddPasswordItem(database *sql.DB, password PasswordItem) error {
	statement := `INSERT INTO password_items (user, app, password, note) VALUES(?,?,?,?) `

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	_, err = database.Exec(statement, password.User, password.App, hashedPassword, password.Note)
	if err != nil {
		return err
	}

	return nil
}

func GetPasswordItems(database *sql.DB) ([]PasswordItem, error) {
	var passwordItems []PasswordItem
	statement := "SELECT * FROM password_items"
	result, err := database.Query(statement)
	defer result.Close()
	if err != nil {
		return passwordItems, err
	}

	for result.Next() {
		var password PasswordItem
		err = result.Scan(&password.Id, &password.User, &password.App, &password.Password, &password.Note)
		if err != nil {
			continue
		}
		passwordItems = append(passwordItems, password)

	}
	return passwordItems, nil

}
