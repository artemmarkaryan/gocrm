package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/order"
)

func (s Service) GetAll() (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var orders []domain.Order
	db.Find(&orders)

	allOrdersPreview := order.ManyOrdersPreview{}
	for _, o := range orders {
		allOrdersPreview.OrderPreviews = append(
			allOrdersPreview.OrderPreviews,
			*order.CreateOrderPreview(&o),
		)
	}

	return allOrdersPreview.Serialize()
}
