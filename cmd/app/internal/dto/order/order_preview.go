package order

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type OrderPreview struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"userId"`
	CustomerID uint   `json:"customerId"`
	Status     string `json:"status"`
}

func CreateOrderPreview(order *domain.Order) *OrderPreview {
	return &OrderPreview{
		ID:         order.ID,
		UserID:     order.UserID,
		CustomerID: order.CustomerID,
		Status:     order.OrderStatus,
	}
}

func (o *OrderPreview) String() string {
	return fmt.Sprintf(`
id: %v
userId: %v
customerId: %v
status: %v`,
		o.ID,
		o.UserID,
		o.CustomerID,
		o.Status,
	)
}

func (o OrderPreview) Deserialize(i interface{}) error {
	panic("implement me")
}

func (o *OrderPreview) Serialize() (string, error) {
	return jsons.MarshalToString(o)
}

func (o *OrderPreview) ToDBO() *domain.Order {
	ord := domain.Order{
		UserID:      o.UserID,
		CustomerID:  o.CustomerID,
		OrderStatus: o.Status,
	}
	ord.ID = o.ID
	return &ord
}
