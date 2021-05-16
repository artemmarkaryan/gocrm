package domain

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	User      User
	Customer  Customer
	CreatedTS time.Time
	Status    OrderStatus
}
