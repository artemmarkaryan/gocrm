package server

import (
	"github.com/gin-gonic/gin"
)

func getEngine() (r *gin.Engine) {
	r = gin.Default()
	setCommonMiddleware(r)
	setRouter(r)

	return
}
