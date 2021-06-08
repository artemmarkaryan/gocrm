package item

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/pkg/jsons"
	"gorm.io/gorm"
)

type Item struct {
	ID        uint  `json:"id"`
	ProductID uint  `json:"productId"`
	BasketID  uint  `json:"basketId"`
	Quantity  uint8 `json:"quantity"`
}

func CreateItem(dbo *domain.Item) *Item {
	return &Item{
		ID:        dbo.ID,
		ProductID: dbo.ProductID,
		BasketID:  dbo.BasketID,
		Quantity:  dbo.Quantity,
	}
}

func (i *Item) ToDBO() *domain.Item {
	return &domain.Item{
		Model:     gorm.Model{
			ID: i.ID,
		},
		BasketID:  i.BasketID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
	}
}

func (i *Item) Serialize() (string, error) {
	return jsons.MarshalToString(i)
}

