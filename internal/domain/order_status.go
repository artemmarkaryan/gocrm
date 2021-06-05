package domain

import "gorm.io/gorm"

type OrderStatus struct {
	gorm.Model
	Name string
}
