package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (view View) New(
	ctx *gin.Context,
) {
	v, err := order.Service{}.New(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"orderId": *v})
	}
}
