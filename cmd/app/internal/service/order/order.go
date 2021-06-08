package order

import (
	"errors"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/order"
	"github.com/gin-gonic/gin"
	"log"
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
			*order.CreateOrderPreview(&o),
		)
	}

	return allOrdersPreview.Serialize()
}

func (s Service) GetOne(ctx *gin.Context) (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var dbOrder domain.Order
	db.Preload("Basket.Items").Find(&dbOrder, ctx.Param("id"))
	return order.CreateOrder(&dbOrder).Serialize()
}

func (s Service) PatchOne(ctx *gin.Context) (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	dbOrder := &domain.Order{}
	db.Find(&dbOrder, ctx.Param("id"))

	dtOrderPreview := order.OrderPreview{}
	dtOrderPreview = *order.CreateOrderPreview(dbOrder)

	err = ctx.BindJSON(&dtOrderPreview)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbOrder = dtOrderPreview.ToDBO()
	db.Save(&dbOrder).Commit()
	return
}

func (s Service) Delete(id string) error {
	db, err := domain.GetDB()
	if err != nil {
		return err
	}

	o := domain.Order{}
	db.Find(&o, id)

	if o.ID == 0 {
		return errors.New("Orders does not exist")
	}
	db.Unscoped().Delete(&o).Commit()

	return nil
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
