package domain

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey;autoIncrement:true"`
	Order Order
}
