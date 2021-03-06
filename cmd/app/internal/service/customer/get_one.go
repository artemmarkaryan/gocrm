package customer

import (
	"errors"
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/customer"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service"
)

func (r Service) GetOne(id string) (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var dbCustomer domain.Customer
	db.Find(&dbCustomer, id)

	if dbCustomer.ID == 0 {
		return "", errors.New(fmt.Sprintf(service.NotFoundF, "customer"))
	}

	return customer.CreateCustomer(dbCustomer).Serialize()
}
