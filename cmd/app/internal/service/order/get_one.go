package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/order"
	"github.com/gin-gonic/gin"
)

func (s Service) GetOne(ctx *gin.Context) (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var dbOrder domain.Order
	db.Preload("Basket.Items").Find(&dbOrder, ctx.Param("id"))
	return order.CreateOrder(&dbOrder).Serialize()
}

