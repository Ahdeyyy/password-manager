package main

import (
	"fmt"
	"log"
	"password-manager/db"
)

func main() {
	store := db.OpenDatabase(":memory:")
	p := db.Profile{Name: "ahdeyy", MasterPassword: "password"}
	err := db.AddProfile(store, p)
	if err != nil {
		log.Printf("%s", err)
	}

	p = db.Profile{Name: "ahdeyy", MasterPassword: "password"}
	err = db.AddProfile(store, p)
	if err != nil {
		log.Printf("%s", err)
	}
	profiles, e := db.GetProfiles(store)
	if e != nil {
		log.Printf("%s", e)
	}
	fmt.Printf("%v", profiles)
}

//  TODO: create profile for storing passwords -> name, master password
//  TODO: Item(Password) -> app/url, user,  password, note
//  TODO: create profile, edit profile, delete profile
//  TODO: create password item, edit item, delete item
