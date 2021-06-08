package basket

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/item"
)

type BasketPreview struct {
	Items []*item.ItemPreview `json:"items"`
}

func CreateBasketPreview(b *domain.Basket) *BasketPreview {
	n := &BasketPreview{}
	for _, i := range b.Items {
		n.Items = append(n.Items, item.CreateItemPreview(&i))
	}
	return n
}
