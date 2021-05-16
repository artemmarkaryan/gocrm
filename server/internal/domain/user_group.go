package domain

import (
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	ID uint64 `gorm:"primaryKey;autoIncrement:true"`
}
