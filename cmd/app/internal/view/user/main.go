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
	data := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := ctx.BindJSON(&data)
	if err != nil || data.Username == "" || data.Password == "" {
		ctx.Status(http.StatusBadRequest)
	}

	ok, _ := user.Service{}.Auth(data.Username, data.Password)
	if ok {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}