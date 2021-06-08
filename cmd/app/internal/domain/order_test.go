package domain

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"testing"
)

func TestOrder(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	db, err := GetDB()
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	db.Create(&Order{
		Basket: Basket{
			Items: []Item{
				{
					ProductID: 1,
					Quantity:  1,
				},
			},
		},
		UserID:      4,
		CustomerID:  3,
		OrderStatus: "Новый",
	}).Commit()
}
