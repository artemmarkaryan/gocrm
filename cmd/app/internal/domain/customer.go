package domain

import (
	"github.com/artemmarkaryan/gocrm/src/pkg/jsonb"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Contact jsonb.JSONB
	Name    string
	Orders  []Order // Customer has many Order
}