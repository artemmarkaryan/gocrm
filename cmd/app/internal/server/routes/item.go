package routes

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/item"
	"github.com/gin-gonic/gin"
)

type Item struct {}

func (i Item) AddRoutes(r *gin.Engine) {
	r.PATCH("/item/:id", item.View{}.PatchOne)
	r.DELETE("/item/:id", item.View{}.Delete)
}


