package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/customer"
	"github.com/gin-gonic/gin"
	"net/http"
)

type View struct{}

func (view View) GetAll(
	ctx *gin.Context,
) {
	v, _ := customer.Service{}.GetAll()
	ctx.String(http.StatusOK, v)
}

func (view View) GetOne(
	ctx *gin.Context,
) {
	id := ctx.Param("id")
	v, _ := customer.Service{}.GetOne(id)
	ctx.String(http.StatusOK, v)
}
