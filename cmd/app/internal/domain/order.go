package domain

import (
	"fmt"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// Order has one Basket
	Basket Basket `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// User has many Order
	User   User
	UserID uint

	// Customer has many Order
	Customer   Customer
	CustomerID uint

	OrderStatus string
}

func (o *Order) String() string {
	return fmt.Sprintf(`
id: %v,
user: %v,
customer: %v,
status: %v,
basket: {%v}
`,
		o.ID,
		o.UserID,
		o.CustomerID,
		o.OrderStatus,
		o.Basket,
	)
}
