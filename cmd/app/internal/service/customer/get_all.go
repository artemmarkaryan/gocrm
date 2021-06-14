package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/customer"
)

func (r Service) GetAll() (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var customers []domain.Customer
	db.Find(&customers)

	allCustomersPreview := customer.ManyCustomersPreview{}
	for _, c := range customers {
		allCustomersPreview.CustomersPreview = append(
			allCustomersPreview.CustomersPreview,
			*customer.CreateCustomerPreview(c),
		)
	}

	return allCustomersPreview.Serialize()
}

