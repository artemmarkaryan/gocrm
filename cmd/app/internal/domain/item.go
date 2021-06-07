package domain

import (
	"fmt"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model

	BasketId uint // Basket has many Item

	Product   Product // Item belongs to Product
	ProductID uint

	Quantity uint8
}

func (i Item) String() string {
	return fmt.Sprintf(`
productId: %v,
quantity: %v
`,
		i.ProductID,
		i.Quantity)
}
