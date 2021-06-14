package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsonb"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
	"gorm.io/gorm"
)

type Customer struct {
	ID      uint         `json:"id"`
	Name    string       `json:"name"`
	Contact *jsonb.JSONB `json:"contact"`
}

func CreateCustomer(dbo domain.Customer) *Customer {
	return &Customer{
		ID:      dbo.ID,
		Name:    dbo.Name,
		Contact: dbo.Contact,
	}
}

func (c *Customer) Serialize() (string, error) {
	return jsons.MarshalToString(c)
}

func (c *Customer) ToDBO() *domain.Customer {
	return &domain.Customer{
		Model:   gorm.Model{ID: c.ID},
		Name:    c.Name,
		Contact: c.Contact,
	}
}