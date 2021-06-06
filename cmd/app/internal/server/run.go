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

	log.Printf("Running on http://%v", address)
	_ = r.Run(address)
}
