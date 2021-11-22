package entities

type User struct {
	UserID                int32
	pwdHash               string
	Email                 string
	Name                  string
	Age                   string
	AdditionalInformation string
	Parent                []User
}
