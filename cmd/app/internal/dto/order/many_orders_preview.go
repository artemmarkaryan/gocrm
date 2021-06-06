package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
)

type ManyOrdersPreview struct {
	OrderPreviews []OrderPreview
}

func (r ManyOrdersPreview) Serialize() (result string, err error, ) {
	return jsons.MarshalToString(r.OrderPreviews)
}
