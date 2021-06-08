package product

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/product"
)

type Service struct{}

func (userService Service) GetAll() (result string, err error) {
	db, err := domain.GetDB()
	if err != nil {
		return
	}

	var products []domain.Product
	db.Find(&products)

	allProductsPreview := product.ManyProducts{}
	for _, p := range products {
		allProductsPreview.Products = append(
			allProductsPreview.Products,
			*product.CreateProduct(p),
		)
	}

	return allProductsPreview.Serialize()
}
