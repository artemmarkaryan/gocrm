package server

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/customer"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/order"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/product"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/user"
	"github.com/gin-gonic/gin"
	"log"
)

const address = "0.0.0.0:8000"

func Run() {
	r := gin.Default()
	r.GET("/user/all", user.View{}.GetAll)
	r.GET("/customer/all", customer.View{}.GetAll)
	r.GET("/product/all", product.View{}.GetAll)
	r.GET("/order/all", order.View{}.GetAll)
	r.GET("/order/:id", order.View{}.GetOne)
	r.PATCH("/order/:id", order.View{}.PatchOne)
	r.DELETE("/order/:id", order.View{}.Delete)
	r.POST("/order", order.View{}.New)

	log.Printf("Running on http://%v", address)
	_ = r.Run(address)
}
