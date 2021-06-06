package user

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type View struct{}

func (view View) GetAll(
	ctx *gin.Context,
) {
	v, _ := user.Service{}.GetAll()
	ctx.String(http.StatusOK, v)
}
