package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type View struct {}

func (view View) Get(
	ctx *gin.Context,
) {
	ctx.Status(http.StatusOK)
}