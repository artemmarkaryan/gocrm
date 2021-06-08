package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string

	// User has many Order
	Orders []Order
}
