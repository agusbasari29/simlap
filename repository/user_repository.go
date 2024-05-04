package repository

import (
	"github.com/agusbasari29/simlap.git/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user entity.Users) (entity.Users, error)
	UpdateUser(user entity.Users) (entity.Users, error)
	GetUser(user entity.Users) (entity.Users, error)
	GetUserByUsername(username string) interface{}
	GetUsers(users []entity.Users) ([]entity.Users, error)
	DeleteUser(user entity.Users) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user entity.Users) (entity.Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user entity.Users) (entity.Users, error) {
	err := r.db.Model(&user).Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUser(user entity.Users) (entity.Users, error) {
	err := r.db.Where("username = ?", user.Username).Take(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) interface{} {
	user := entity.Users{}
	result := r.db.Where("username = ?", username).Take(&user)
	if result.Error == nil {
		return result
	}
	return nil
}

func (r *userRepository) GetUsers() ([]entity.Users, error) {
	users := []entity.Users{}
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepository) DeleteUser(user entity.Users) bool {
	return r.db.Delete(&user).Error == nil
}
