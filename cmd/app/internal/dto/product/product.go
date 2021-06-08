package product

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type Product struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func (p Product) Deserialize(i interface{}) error {
	panic("implement me")
}

func (p Product) Serialize() (string, error) {
	return jsons.MarshalToString(p)
}

func CreateProduct(dbo domain.Product) *Product {
	return &Product{Id: dbo.ID, Name: dbo.Name, Price: dbo.Price}
}
