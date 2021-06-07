package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsonb"
)

type CustomerPreview struct {
	ID      uint    `json:"id"`
	Name    string  `json:"name"`
	Contact *jsonb.JSONB `json:"contact"`
}

func (c CustomerPreview) Serialize() (string, error) {
	panic("implement me")
}

func (c CustomerPreview) Deserialize(i interface{}) error {
	panic("implement me")
}

func NewCustomerPreview(dbo domain.Customer) *CustomerPreview {
	return &CustomerPreview{ID: dbo.ID, Name: dbo.Name, Contact: dbo.Contact}
}
