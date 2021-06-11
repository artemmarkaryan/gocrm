package server

import (
	"github.com/gin-gonic/gin"
)

func getEngine() (r *gin.Engine) {
	r = gin.Default()
	setCommonMiddleware(r)

	//r.POST("/login", user.View{}.Auth)
	setRouter(r)

	return
}
