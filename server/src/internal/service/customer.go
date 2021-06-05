package service

import (
	"github.com/artemmarkaryan/gocrm/internal/domain"
	customerDTO "github.com/artemmarkaryan/gocrm/internal/dto/customer"
)

type CustomerService struct {}

func (r CustomerService) GetAll() (result []byte, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var customers []domain.Customer
	db.Find(&customers)

	return customerDTO.AllCustomerPreview{}.Serialize(customers)
}

