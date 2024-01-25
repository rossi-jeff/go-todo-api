package dto

type LoginCredentials struct {
	UserName string
	PassWord string
}

type RegisterCredentials struct {
	LoginCredentials
	Email string
}

type LoginResponse struct {
	UserName string
	Token    string
}
