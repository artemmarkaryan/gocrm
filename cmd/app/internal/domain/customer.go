package domain

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsonb"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Contact *jsonb.JSONB
	Name    string

	// Customer has many Order
	Orders  []Order
}
