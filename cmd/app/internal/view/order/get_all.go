package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (view View) GetAll(
	ctx *gin.Context,
) {
	v, _ := order.Service{}.GetAll()
	ctx.String(http.StatusOK, v)
}

