package server

import "github.com/gin-gonic/gin"

func getEngine() *gin.Engine {
	return gin.Default()
}
