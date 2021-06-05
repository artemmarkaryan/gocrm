package domain

import "gorm.io/gorm"

type Basket struct {
	gorm.Model

	OrderID uint // Order has one Basket

	Items []Item // Basket has many Item
}
