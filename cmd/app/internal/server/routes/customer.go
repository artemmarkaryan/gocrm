package routes

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/customer"
	"github.com/gin-gonic/gin"
)

type Customer struct{}

func (c Customer) AddRoutes(r *gin.Engine) {
	r.GET("/customer/all", customer.View{}.GetAll)
	r.GET("/customer/:id", customer.View{}.GetOne)

}
