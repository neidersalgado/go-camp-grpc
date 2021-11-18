package entities

type User struct {
	Id                    string
	PwdHash               string
	EMail                 string
	Name                  string
	Age                   int32
	AdditionalInformation string
	Parents               []User
}
