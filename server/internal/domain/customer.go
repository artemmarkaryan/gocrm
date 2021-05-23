package domain

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Contact pgtype.JSONB

	Orders  []Order // Customer has many Order
}
