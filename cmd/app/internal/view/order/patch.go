package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (view View) PatchOne(
	ctx *gin.Context,
) {
	_, err := order.Service{}.PatchOne(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.Status(http.StatusOK)
	}
}
