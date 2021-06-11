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

func (view View) Auth(
	ctx *gin.Context,
) {
	username, password := ctx.Param("username"), ctx.Param("password")
	if username == "" || password == "" {
		ctx.Status(http.StatusBadRequest)
	}
	ok, _ := user.Service{}.Auth(username, password)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
	} else {

	}
}