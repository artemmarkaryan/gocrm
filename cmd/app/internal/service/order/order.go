package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/order"
	"github.com/gin-gonic/gin"
)

type Service struct{}

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
			*order.NewOrderPreview(o),
		)
	}

	return allOrdersPreview.Serialize()
}

func (s Service) New(ctx *gin.Context) (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	dtOrder := order.NewOrder{}
	err = ctx.BindJSON(&dtOrder)
	if err != nil {
		return
	}
	dbOrder := dtOrder.ToDBO()
	db.Create(dbOrder).Commit()
	return
}
