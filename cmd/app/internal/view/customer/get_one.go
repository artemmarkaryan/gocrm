package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/customer"
	"github.com/gin-gonic/gin"
	"net/http"
)



func (r View) GetOne(
	ctx *gin.Context,
) {
	id := ctx.Param("id")
	v, _ := customer.Service{}.GetOne(id)
	ctx.String(http.StatusOK, v)
}
