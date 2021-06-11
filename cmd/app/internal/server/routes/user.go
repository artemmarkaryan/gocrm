package routes

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/user"
	"github.com/gin-gonic/gin"
)

type User struct {}

func (u User) AddRoutes(r *gin.Engine) {
	r.GET("/user/all", user.View{}.GetAll)
	r.POST("/auth", user.View{}.Auth)
}

