package routes

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/user"
	"github.com/gin-gonic/gin"
)

type User struct {}

func (u User) AddRoutes(r *gin.RouterGroup) {
	r.GET("/user/all", user.View{}.GetAll)

}

