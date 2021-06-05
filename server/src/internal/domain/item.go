package domain

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model

	BasketId uint // Basket has many Item

	Product   Product // Item belongs to Product
	ProductID uint

	Quantity uint8
}
