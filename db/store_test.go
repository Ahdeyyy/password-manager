package db

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestOpenDatabase(t *testing.T) {
	db := OpenDatabase(":memory:")
	if db == nil {
		t.Errorf("failed to open database")
	}
	defer db.Close()
}

func TestProfileCreateAndGet(t *testing.T) {
	db := OpenDatabase(":memory:")
	p := Profile{Name: "ahdeyy", MasterPassword: "password"}
	p2 := Profile{Name: "ahdeyy2", MasterPassword: "password123"}
	err := AddProfile(db, p)
	if err != nil {
		t.Errorf("failed to add profile: %s", err)
	}

	err = AddProfile(db, p2)
	if err != nil {
		t.Errorf("failed to add profile: %s", err)
	}

	profiles, err := GetProfiles(db)
	if err != nil {
		t.Errorf("failed to get profiles: %s", err)
	}

	if len(profiles) != 2 {
		t.Errorf("expected 2 profiles, got %d", len(profiles))
	}

	if profiles[0].Name != p.Name {
		t.Errorf("expected profile name '%s', got '%s'", p.Name, profiles[0].Name)
	}

	if profiles[1].Name != p2.Name {
		t.Errorf("expected profile name '%s', got '%s'", p2.Name, profiles[1].Name)
	}

	if bcrypt.CompareHashAndPassword([]byte(profiles[0].MasterPassword), []byte(p.MasterPassword)) != nil {
		fmt.Printf("%s", profiles[0].MasterPassword)
		t.Errorf("Hashed passwords do not match")
	}

	if bcrypt.CompareHashAndPassword([]byte(profiles[1].MasterPassword), []byte(p2.MasterPassword)) != nil {
		t.Errorf("Hashed passwords do not match")
	}

}

func TestProfileEditAndDelete(t *testing.T) {
	db := OpenDatabase(":memory:")
	p := Profile{Name: "ahdeyy", MasterPassword: "password"}
	_ = AddProfile(db, p)
	profiles, _ := GetProfiles(db)

	new_p := profiles[0]

	new_p.MasterPassword = "password"
	new_p.Name = "ahdeyy2"

	err := EditProfile(db, new_p)

	if err != nil {
		t.Errorf("failed to edit profile: %s", err)
	}

	profile2, _ := GetProfile(db, int(new_p.Id))

	if profile2.Name != new_p.Name {
		t.Errorf("expected profile name '%s', got '%s'", new_p.Name, profile2.Name)
	}

	if bcrypt.CompareHashAndPassword([]byte(profile2.MasterPassword), []byte(new_p.MasterPassword)) != nil {
		fmt.Printf("%s", new_p.MasterPassword)
		t.Errorf("Hashed passwords do not match")
	}

	err = DeleteProfile(db, new_p)

	if err != nil {
		t.Errorf("failed to delete profile: %s", err)
	}

	profiles3, _ := GetProfiles(db)

	if len(profiles3) != 0 {
		t.Errorf("expected 0 profiles, got %d", len(profiles3))
	}

}

func TestCreateAndGetPasswordItems(t *testing.T) {

	db := OpenDatabase(":memory:")
	p := PasswordItem{User: "ahdeyy", App: "google", Password: "password", Note: "test"}
	err := AddPasswordItem(db, p)
	if err != nil {
		t.Errorf("failed to add password item: %s", err)
	}
	p2 := PasswordItem{User: "ahdeyy", App: "facebook", Password: "password", Note: "test"}

	err = AddPasswordItem(db, p2)

	if err != nil {
		t.Errorf("failed to add password item: %s", err)
	}

	items, err := GetPasswordItems(db)
	if err != nil {
		t.Errorf("failed to get password items: %s", err)
	}

	if len(items) != 2 {
		t.Errorf("expected 2 password items, got %d", len(items))
	}

	if items[0].App != p.App {
		t.Errorf("expected app: '%s', got '%s'", p.App, items[0].App)
	}

	if items[1].App != p2.App {
		t.Errorf("expected app: '%s', got '%s'", p.App, items[1].App)
	}

}

func TestEditAndDeletePasswordItem(t *testing.T) {
	db := OpenDatabase(":memory:")
	p := PasswordItem{User: "ahdeyy", App: "google", Password: "password", Note: "test"}
	_ = AddPasswordItem(db, p)
	items, _ := GetPasswordItems(db)
	newP := items[0]

	newP.Note = "test2"

	err := EditPasswordItem(db, newP)

	if err != nil {

		t.Errorf("failed to edit password item: %s", err)
	}

	items2, _ := GetPasswordItems(db)

	if items2[0].Note != newP.Note {
		t.Errorf("expected note: '%s', got '%s'", newP.Note, items2[0].Note)
	}

	err = DeletePasswordItem(db, newP)

	if err != nil {
		t.Errorf("failed to delete password item: %s", err)
	}

	items3, _ := GetPasswordItems(db)

	if len(items3) != 0 {
		t.Errorf("expected 0 password items, got %d", len(items3))
	}

}
