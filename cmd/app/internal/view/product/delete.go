package product

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (view View) Delete(ctx *gin.Context) {
	err := product.Service{}.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.Status(http.StatusOK)
	}
}

