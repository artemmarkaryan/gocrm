package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (view View) Delete(
	ctx *gin.Context,
) {
	err := order.Service{}.Delete(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
	} else {
		ctx.Status(http.StatusOK)
	}
}

