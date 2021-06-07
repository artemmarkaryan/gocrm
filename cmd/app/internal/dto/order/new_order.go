package order

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/item"
)

type NewOrder struct {
	UserId     uint               `json:"userId"`
	CustomerId uint               `json:"customerId"`
	Items      []item.ItemPreview `json:"items"`
}

func (o NewOrder) String() string {
	return fmt.Sprintf(`
userId: %v
customerId: %v, 
items: %v`,
		o.UserId,
		o.CustomerId,
		o.Items)
}

func (n NewOrder) Serialize() (string, error) {
	panic("implement me")
}

func (o NewOrder) ToDBO() *domain.Order {

	var dbItems []domain.Item

	for _, i := range o.Items {
		dbItems = append(dbItems, *i.ToDBO())
	}

	return &domain.Order{
		Basket: domain.Basket{
			Items: dbItems,
		},
		UserID:        o.UserId,
		CustomerID:    o.CustomerId,
		OrderStatusID: 1,
	}
}
