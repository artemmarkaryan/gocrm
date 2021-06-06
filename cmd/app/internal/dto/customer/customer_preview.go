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

func NewCustomerPreview(dbo domain.Customer) *CustomerPreview {
	return &CustomerPreview{ID: dbo.ID, Name: dbo.Name, Contact: dbo.Contact}
}
