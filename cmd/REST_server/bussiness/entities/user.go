package entities

import "time"

type User struct {
	Id                    string
	PwdHash               string
	EMail                 string
	Name                  string
	BirthDate             time.Time
	AdditionalInformation string
	Parents               []User
}
