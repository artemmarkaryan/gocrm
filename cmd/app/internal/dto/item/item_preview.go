package item

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
)

type ItemPreview struct {
	Id        uint  `json:"id"`
	ProductId uint  `json:"productId"`
	Quantity  uint8 `json:"quantity"`
}

func CreateItemPreview(i *domain.Item) *ItemPreview {
	return &ItemPreview{Id: i.ID, ProductId: i.ProductID, Quantity: i.Quantity}
}

func (i ItemPreview) String() string {
	return fmt.Sprintf(`
id: %v,
productId: %v,
quantity: %v`,
		i.Id,
		i.ProductId,
		i.Quantity)
}

func (i ItemPreview) Serialize() (string, error) {
	panic("implement me")
}

func (i ItemPreview) Deserialize(i2 interface{}) error {
	panic("implement me")
}

func (i ItemPreview) ToDBO() *domain.Item {
	return &domain.Item{
		ProductID: i.ProductId,
		Quantity:  i.Quantity,
	}
}
