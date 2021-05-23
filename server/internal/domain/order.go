package domain

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Basket Basket // Order has one Basket

	User   User // User has many Order
	UserID uint

	Customer   Customer // Customer has many Order
	CustomerID uint

	OrderStatus   OrderStatus // Order belongs to OrderStatus
	OrderStatusID uint
}
