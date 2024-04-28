package repository

import (
	"github.com/agusbasari29/simlap.git/entity"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	CreateUserRole(userRole entity.UserRole) (entity.UserRole, error)
	UpdateUserRole(userRole entity.UserRole) (entity.UserRole, error)
	GetUserRole(userRole entity.UserRole) (entity.UserRole, error)
	GetUserRoles(usersRole []entity.UserRole) ([]entity.UserRole, error)
	DeleteUserRole(userRole entity.UserRole) bool
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *userRoleRepository {
	return &userRoleRepository{db}
}

func (r *userRoleRepository) CreateUserRole(userRole entity.UserRole) (entity.UserRole, error) {
	err := r.db.Create(&userRole).Error
	if err != nil {
		return userRole, err
	}
	return userRole, nil
}

func (r *userRoleRepository) UpdateUserRole(userRole entity.UserRole) (entity.UserRole, error) {
	err := r.db.Model(&userRole).Save(&userRole).Error
	if err != nil {
		return userRole, err
	}
	return userRole, nil
}

func (r *userRoleRepository) GetUserRole(userRole entity.UserRole) (entity.UserRole, error) {
	err := r.db.Where("id = ?", userRole.ID).Take(&userRole).Error
	if err != nil {
		return userRole, err
	}
	return userRole, nil
}

func (r *userRoleRepository) GetUserRoles(userRoles []entity.UserRole) ([]entity.UserRole, error) {
	err := r.db.Find(&userRoles).Error
	if err != nil {
		return userRoles, err
	}
	return userRoles, nil
}

func (r *userRoleRepository) DeleteUserRole(userRole entity.UserRole) bool {
	return r.db.Delete(&userRole).Error == nil
}
