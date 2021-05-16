package domain

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey;autoIncrement:true"`
	Contact pgtype.JSONB
}
