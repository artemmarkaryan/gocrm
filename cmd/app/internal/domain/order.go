package domain

import (
	"fmt"
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

func (o Order) String() string {
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
		o.OrderStatusID,
		o.Basket,
	)
}
