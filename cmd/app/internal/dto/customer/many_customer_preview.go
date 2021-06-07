package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type ManyCustomersPreview struct {
	CustomersPreview []CustomerPreview
}

func (r ManyCustomersPreview) Deserialize(i interface{}) error {
	panic("implement me")
}

func (r ManyCustomersPreview) Serialize() (string, error) {
	return jsons.MarshalToString(r.CustomersPreview)
}
