package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/customer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r View) New(ctx *gin.Context) {
	customerId, err := customer.Service{}.New(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"customerId": customerId})
	}
}
