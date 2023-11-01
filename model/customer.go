package model

type CreateUserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
