package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Login    string
	Password string `gorm:"type:bytea"`

	// User has many Order
	Orders []Order
}
