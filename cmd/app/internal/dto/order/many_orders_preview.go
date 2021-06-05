package order

import (
	"encoding/json"
)

type ManyOrdersPreview struct {
	OrderPreviews []OrderPreview
}

func (r ManyOrdersPreview) Serialize() (result []byte, err error, ) {
	return json.Marshal(r.OrderPreviews)
}
