package models

type User struct {
	Id                    string `json:"Id`
	PwdHash               string `json:"pwdhash"`
	EMail                 string `json:"email"`
	Name                  string `json:"name"`
	Age                   int32  `json:"age"`
	AdditionalInformation string `json:"additional_information"`
	Parents               []User `json:"parents"`
}
