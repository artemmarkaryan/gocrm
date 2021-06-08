package server

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/customer"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/item"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/order"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/product"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/user"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

const address = "0.0.0.0"

func Run() {
	r := gin.Default()

	// user
	r.GET("/user/all", user.View{}.GetAll)

	// customer
	r.GET("/customer/all", customer.View{}.GetAll)
	r.GET("/customer/:id", customer.View{}.GetOne)

	// product
	r.GET("/product/all", product.View{}.GetAll)

	// item
	r.PATCH("/item/:id", item.View{}.PatchOne)
	r.DELETE("/item/:id", item.View{}.Delete)

	// order
	r.GET("/order/all", order.View{}.GetAll)
	r.GET("/order/:id", order.View{}.GetOne)
	r.PATCH("/order/:id", order.View{}.PatchOne)
	r.DELETE("/order/:id", order.View{}.Delete)
	r.POST("/order", order.View{}.New)
	
	portString := os.Args[0]
	port, err := strconv.Atoi(portString)
	if err != nil  {
		log.Fatal("No port declared in env")
	}
	log.Printf("Running on http://%v:%v", address, port)
	_ = r.Run(address)
}
