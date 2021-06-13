package routes

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/user"
	"github.com/gin-gonic/gin"
)

type Auth struct {}

func (a Auth) AddRoutes(r *gin.Engine) {
	r.POST("/login", user.View{}.Auth)
}

