package response

import "github.com/agusbasari29/simlap.git/entity"

type UserRegisterResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     entity.UserRole
}
