package item

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/item"
	"github.com/gin-gonic/gin"
	"net/http"
)

type View struct {}

func (view View) PatchOne(
	ctx *gin.Context,
) {
	v, _ := item.Service{}.PatchOne(ctx)
	ctx.String(http.StatusOK, v)
}

func (view View) Delete(
	ctx *gin.Context,
) {
	err := item.Service{}.Delete(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
	} else {
		ctx.Status(http.StatusOK)
	}
}
