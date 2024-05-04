package request

import "github.com/agusbasari29/simlap.git/entity"

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Role     entity.UserRole
}

type UserGetRequest struct{
	ID uint64 `json:"id"`
}