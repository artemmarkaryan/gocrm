package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/customer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r View) GetAll(ctx *gin.Context) {
	v, _ := customer.Service{}.GetAll()
	ctx.String(http.StatusOK, v)
}
