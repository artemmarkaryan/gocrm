package domain

import (
	"fmt"
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model

	// Order has one Basket
	OrderID uint

	// Basket has many Item
	Items []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
