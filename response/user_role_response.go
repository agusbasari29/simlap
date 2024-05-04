package response

import "github.com/agusbasari29/simlap.git/entity"

type UserRoleResponse struct {
	ID       uint64          `json:"id"`
	RoleName string          `json:"role_name"`
	Section  entity.Sections `json:"section"`
}
