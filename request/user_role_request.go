package request

import "github.com/agusbasari29/simlap.git/entity"

type UserRoleCreateRequest struct {
	RoleName string          `json:"role_name"`
	Section  entity.Sections `json:"section"`
}

type UserRoleRequest struct {
	ID       uint64          `json:"id"`
	RoleName string          `json:"role_name"`
	Secrion  entity.Sections `json:"section"`
}