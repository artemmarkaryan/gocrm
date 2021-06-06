package product

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

type View struct{}

func (view View) GetAll(
	ctx *gin.Context,
) {
	v, _ := product.Service{}.GetAll()
	ctx.String(http.StatusOK, v)
}
