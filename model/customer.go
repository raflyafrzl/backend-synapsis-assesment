package model

type CreateUserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthUserModel struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}
