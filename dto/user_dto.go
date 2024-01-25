package dto

type UserAttrs struct {
	UserName string
	Email    string
}

type ChangePW struct {
	OldPassWord  string
	NewPassWord  string
	Confirmation string
}
