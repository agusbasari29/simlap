package request

import "github.com/agusbasari29/simlap.git/entity"

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	Username string
	Password string
	Fullname string
	Phone    string
	Email    string
	Address  string
	Role     entity.UserRole
}
