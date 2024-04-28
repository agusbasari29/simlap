package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey,autoIncrement"`
	Username string
	Password string
	Fullname string
	Email    string
	Phone    string
	Address  string
	Role     UserRole
}
