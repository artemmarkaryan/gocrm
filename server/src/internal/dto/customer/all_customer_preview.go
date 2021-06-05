package customer

import (
	"encoding/json"
	"github.com/artemmarkaryan/gocrm/internal/domain"
)

type AllCustomerPreview struct {
	CustomersPreview []CustomerPreview
}

func (r AllCustomerPreview) Serialize(
	customers []domain.Customer,
) (result []byte, err error, ) {
	for _, user := range customers {
		r.CustomersPreview = append(
			r.CustomersPreview,
			CustomerPreview{
				ID:   user.ID,
				Name: user.Name,
			})
	}
	return json.Marshal(r.CustomersPreview)
}
