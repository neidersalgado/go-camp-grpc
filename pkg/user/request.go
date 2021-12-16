package user

type UserRequest struct {
	UserID                int32        `json:"user_id"`
	PWDHash               string        `json:"pwd_hash"`
	Email                 string        `json:"email"`
	Name                  string        `json:"name"`
	Age                   string        `json:"age"`
	AdditionalInformation string        `json:"additional_information"`
	Parent                []UserRequest `json:"parent"`
}
