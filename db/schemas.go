package db

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
