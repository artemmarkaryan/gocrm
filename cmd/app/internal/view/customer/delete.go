package customer

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/customer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r View) Delete(ctx *gin.Context) {
	err := customer.Service{}.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.Status(http.StatusOK)
	}
}
