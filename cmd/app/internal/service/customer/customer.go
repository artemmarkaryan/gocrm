package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	customerDTO "github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/customer"
)

type Service struct{}

func (r Service) GetAll() (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var customers []domain.Customer
	db.Find(&customers)

	allCustomersPreview := customerDTO.ManyCustomersPreview{}
	for _, c := range customers {
		allCustomersPreview.CustomersPreview = append(
			allCustomersPreview.CustomersPreview,
			*customerDTO.CreateCustomerPreview(c),
		)
	}

	return allCustomersPreview.Serialize()
}
