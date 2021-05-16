package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name      string
}
