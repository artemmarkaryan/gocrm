package product

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	dto "github.com/artemmarkaryan/gocrm/cmd/app/internal/dto/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s Service) New(ctx *gin.Context) (productId *uint, err error) {
	var productDTO dto.Product
	var productDBO domain.Product
	var db, operation *gorm.DB
	db, err = domain.GetDB()
	if err != nil {
		return
	}

	err = ctx.BindJSON(&productDTO)
	productDBO = *productDTO.ToDBO()

	operation = db.Create(&productDBO).Commit()
	if operation.RowsAffected == 0 {
		err = operation.Error
	}
	return &productDBO.ID, err
}
