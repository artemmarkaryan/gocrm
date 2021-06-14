package routes

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/ping"
	"github.com/gin-gonic/gin"
)

type Ping struct{}

func (p Ping) AddRoutes(r *gin.RouterGroup) {
	r.GET("/ping", ping.View{}.Get)
}
