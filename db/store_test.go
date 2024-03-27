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

	new_p.MasterPassword = "password123"
	new_p.Name = "ahdeyy2"

	err := EditProfile(db, new_p)

	if err != nil {
		t.Errorf("failed to edit profile: %s", err)
	}

	profiles2, _ := GetProfiles(db)

	if len(profiles2) != 1 {
		t.Errorf("expected 1 profile, got %d", len(profiles2))
	}

	if profiles2[0].Name != new_p.Name {
		t.Errorf("expected profile name '%s', got '%s'", new_p.Name, profiles2[0].Name)
	}

	if bcrypt.CompareHashAndPassword([]byte(profiles2[0].MasterPassword), []byte(new_p.MasterPassword)) != nil {
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
