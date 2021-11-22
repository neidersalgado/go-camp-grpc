package user

type UserModel struct {
	UserID                int32
	PWDHash               string
	Email                 string
	Name                  string
	Age                   string
	AdditionalInformation string
	Parent                []UserModel
}
