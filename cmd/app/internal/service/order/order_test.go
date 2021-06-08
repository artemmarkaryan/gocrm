package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"log"
	"strconv"
	"testing"
)

func TestService_Delete(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	db, err := domain.GetDB()
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	o := domain.Order{
		Basket: domain.Basket{
			Items: []domain.Item{
				{
					ProductID: 1,
					Quantity:  1,
				},
			},
		},
		UserID:      4,
		CustomerID:  3,
		OrderStatus: "Новый",
	}
	db.Create(&o).Commit()

	log.Print("Order created: ", o)

	_, err = Service{}.Delete(strconv.Itoa(int(o.ID)))
}
