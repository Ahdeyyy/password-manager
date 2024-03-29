package db

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+[]{}|;:,.<>?"
)

func GenerateStrongPassword() string {

	b := make([]byte, 32)

	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

// Compares two profiles p1 and p2 and edits p1 with the values of p2 if p1 and p2 are not equal
// p1 calls the function and is mutated
func (p *Profile) CompareAndEdit(profile Profile) {
	if p.Name != profile.Name {
		p.Name = profile.Name
	}

	if bcrypt.CompareHashAndPassword([]byte(p.MasterPassword), []byte(profile.MasterPassword)) != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(profile.MasterPassword), bcrypt.MinCost)
		p.MasterPassword = string(hashedPassword)

	}
}

// compares two passwordItems p1 and p2 and edits p1 with the values of p2 if p1 and p2 are not equal
// p1 calls the function and is mutated

func (p *PasswordItem) CompareAndEdit(password PasswordItem) {
	if p.Note != password.Note {
		p.Note = password.Note
	}

	if bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password.Password)) != nil {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password.Password), bcrypt.MinCost)
		p.Password = string(hashedPassword)
	}

	if p.User != password.User {
		p.User = password.User
	}

	if p.App != password.App {
		p.App = password.App
	}

}
