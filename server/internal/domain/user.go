package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Orders []Order // User has many Order
}
