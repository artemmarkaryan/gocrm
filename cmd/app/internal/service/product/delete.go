package product

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/domain"
	"github.com/gin-gonic/gin"
)


func (s Service) Delete(ctx *gin.Context) (err error) {
	var productDBO domain.Product

	db, err := domain.GetDB()
	if err != nil {
		return
	}

	operation := db.Unscoped().Delete(&productDBO, ctx.Param("id")).Commit()
	if operation.RowsAffected == 0 {
		err = operation.Error
	}

	return
}
