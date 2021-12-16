package user

type UserModel struct {
	UserID                int32
	PWDHash               string
	Email                 string
	Name                  string
	Age                   int32
	AdditionalInformation string
	Parent                []UserModel
}
