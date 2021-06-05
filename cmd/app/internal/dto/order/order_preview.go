package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type OrderPreview struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"userId"`
	CustomerID uint   `json:"customerId"`
	Status     string `json:"status"`
}

func NewOrderPreview(order domain.Order) *OrderPreview {
	return &OrderPreview{
		ID:         order.ID,
		UserID:     order.UserID,
		CustomerID: order.CustomerID,
		Status:     order.OrderStatus.Name,
	}
}

func (o OrderPreview) Serialize() (string, error) {
	return jsons.MarshallToString(o)
}


