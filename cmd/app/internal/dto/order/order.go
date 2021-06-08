package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/basket"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type Order struct {
	Id         uint                  `json:"id"`
	UserId     uint                  `json:"userId"`
	CustomerId uint                  `json:"customerId"`
	Basket     *basket.BasketPreview `json:"basket"`
	Status     string                `json:"status"`
}

func CreateOrder(o *domain.Order) *Order {
	return &Order{
		Id:         o.ID,
		UserId:     o.UserID,
		CustomerId: o.CustomerID,
		Status:     o.OrderStatus,
		Basket:     basket.CreateBasketPreview(&o.Basket),
	}
}

func (o *Order) Serialize() (string, error) {
	return jsons.MarshalToString(o)
}
