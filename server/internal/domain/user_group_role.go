package domain

import "gorm.io/gorm"

type UserGroupRole struct {
	gorm.Model
	Level uint8
	Name  string
}
