package routes

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/order"
	"github.com/gin-gonic/gin"
)

type Order struct {}

func (o Order) AddRoutes(r *gin.RouterGroup) {
	r.GET("/order/all", order.View{}.GetAll)
	r.GET("/order/:id", order.View{}.GetOne)
	r.PATCH("/order/:id", order.View{}.PatchOne)
	r.DELETE("/order/:id", order.View{}.Delete)
	r.POST("/order", order.View{}.New)
}

