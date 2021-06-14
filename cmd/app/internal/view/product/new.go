package product

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (view View) New(
	ctx *gin.Context,
) {
	productId, err := product.Service{}.New(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"productId": productId})
	}
}

