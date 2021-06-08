package domain

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	db, err := GetDB()
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	product := Product{}
	db.Last(&product)

	user := User{}
	db.Last(&user)

	customer := Customer{}
	db.Last(&customer)

	db.Create(&Order{
		Basket: Basket{
			Items: []Item{
				{
					ProductID: product.ID,
					Quantity:  1,
				},
			},
		},
		UserID:      user.ID,
		CustomerID:  customer.ID,
		OrderStatus: "Новый",
	}).Commit()
}

func TestDeleteOrder(t *testing.T) {
	if err := internal.Setup(); err != nil {
		t.Fatal(err.Error())
	}

	db, err := GetDB()
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	product := Product{}
	db.Last(&product)

	user := User{}
	db.Last(&user)

	customer := Customer{}
	db.Last(&customer)

	o := Order{
		Basket: Basket{
			Items: []Item{
				{
					ProductID: product.ID,
					Quantity:  1,
				},
			},
		},
		UserID:      user.ID,
		CustomerID:  customer.ID,
		OrderStatus: "Новый",
	}

	db.Create(&o).Commit()
	db.Unscoped().Delete(&o).Commit()
}
