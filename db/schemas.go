package db

import "golang.org/x/crypto/bcrypt"

type Profile struct {
	Id             int64
	Name           string
	MasterPassword string
}

type PasswordItem struct {
	Id       int64
	User     string
	App      string // can also be a url
	Password string
	Note     string
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
