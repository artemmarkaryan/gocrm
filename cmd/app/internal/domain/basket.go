package domain

import (
	"fmt"
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model

	OrderID uint // Order has one Basket

	Items []Item // Basket has many Item
}

func (b Basket) String() string {
	return fmt.Sprintf(`
id: %v
items: %v
`,
		b.ID,
		b.Items,
	)
}
