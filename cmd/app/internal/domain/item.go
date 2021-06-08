package domain

import (
	"fmt"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model

	// Basket has many Item
	BasketID uint

	// Item belongs to Product
	Product   Product
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
