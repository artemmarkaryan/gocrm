package product

import "github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"

type ManyProducts struct {
	Products []Product
}

func (m ManyProducts) Deserialize(i interface{}) error {
	panic("implement me")
}

func (m ManyProducts) Serialize() (string, error) {
	return jsons.MarshalToString(m.Products)
}
