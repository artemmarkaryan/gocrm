package server

import (
	"fmt"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/customer"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/item"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/order"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/product"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

const address = "0.0.0.0"

func Run() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://0.0.0.0", "http://127.0.0.1", "http://localhost"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE"},
		AllowHeaders:     []string{"X-Requested-With", "content-type", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

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

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("No port declared in env")
	}
	log.Printf("Running on http://%v:%v", address, port)
	_ = r.Run(fmt.Sprintf("%v:%v", address, port))
}
