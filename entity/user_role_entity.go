package entity

import "gorm.io/gorm"

type Sections string

const (
	Dinas   Sections = "dinas"
	Pelapor Sections = "pelapor"
)

type UserRole struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	RoleName string
	Section  Sections
}
