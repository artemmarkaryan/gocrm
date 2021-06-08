package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
)

type CustomerPreview struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (c CustomerPreview) Serialize() (string, error) {
	panic("implement me")
}

func (c CustomerPreview) Deserialize(i interface{}) error {
	panic("implement me")
}

func CreateCustomerPreview(dbo domain.Customer) *CustomerPreview {
	return &CustomerPreview{ID: dbo.ID, Name: dbo.Name}
}
