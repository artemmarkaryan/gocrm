package order

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

type View struct{}

func (view View) GetAll(
	ctx *gin.Context,
) {
	v, _ := order.Service{}.GetAll()
	ctx.String(http.StatusOK, v)
}

func (view View) GetOne(
	ctx *gin.Context,
) {
	v, _ := order.Service{}.GetOne(ctx)
	ctx.String(http.StatusOK, v)
}

func (view View) PatchOne(
	ctx *gin.Context,
) {
	v, _ := order.Service{}.PatchOne(ctx)
	ctx.String(http.StatusOK, v)
}

func (view View) Delete(
	ctx *gin.Context,
) {
	v, _ := order.Service{}.Delete(ctx)
	ctx.String(http.StatusOK, v)
}

func (view View) New(
	ctx *gin.Context,
) {
	v, _ := order.Service{}.New(ctx)
	ctx.String(http.StatusOK, v)
}
