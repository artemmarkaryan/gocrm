package domain

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement:true"`
	Basket   Basket
	Product  Product
	Quantity uint8
}
