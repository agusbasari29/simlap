package service

import (
	"log"

	"github.com/agusbasari29/simlap.git/entity"
	"github.com/agusbasari29/simlap.git/repository"
	"github.com/agusbasari29/simlap.git/request"
	"github.com/mashingan/smapping"
)

type UserRoleService interface {
}

type userRoleService struct {
	userRoleRepository repository.UserRoleRepository
}

func NewUserRoleService(userRoleRepository repository.UserRoleRepository) *userRoleService {
	return &userRoleService{userRoleRepository}
}

func (s *userRoleService) CreateUserRole(request request.UserRoleCreateRequest) (entity.UserRole, error) {
	role := entity.UserRole{}
	role.RoleName = request.RoleName
	if request.Section == "dinas" {
		role.Section = entity.Dinas
	} else {
		role.Section = entity.Pelapor
	}
	newRole, err := s.userRoleRepository.CreateUserRole(role)
	if err != nil {
		return newRole, err
	}
	return newRole, err
}

func (s *userRoleService) UpdateUserRole(request request.UserRoleRequest) (entity.UserRole, error) {
	role := entity.UserRole{}
	err := smapping.FillStruct(&role, smapping.MapFields(&request))
	if err != nil {
		log.Fatalf("error mapping struct: %v", err)
	}
	newRole, err := s.userRoleRepository.UpdateUserRole(role)
	if err != nil {
		return newRole, err
	}
	return newRole, nil
}

func (s *userRoleService) GetUserRole(request request.UserGetRequest) (entity.UserRole, error) {
	role := entity.UserRole{}
	err := smapping.FillStruct(&role, smapping.MapFields(&request))
	if err != nil {
		log.Fatalf("error mapping struct: %v", err)
	}
	getRole, err := s.userRoleRepository.GetUserRole(role)
	if err != nil {
		return getRole, err
	}
	return getRole, nil
}

func (s *userRoleService) GetUserRoles() ([]entity.UserRole, error) {
	roles := []entity.UserRole{}
	getRoles, err := s.userRoleRepository.GetUserRoles(roles)
	if err != nil {
		return getRoles, err
	}
	return getRoles, nil
}

func (s *userRoleService) DeleteUserRole(request request.UserGetRequest) bool {
	role := entity.UserRole{}
	role.ID = request.ID
	return s.userRoleRepository.DeleteUserRole(role)
}
