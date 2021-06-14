package routes

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/view/product"
	"github.com/gin-gonic/gin"
)

type Product struct {}

func (p Product) AddRoutes(r *gin.RouterGroup) {
	r.GET("/product/all", product.View{}.GetAll)
	r.POST("/product", product.View{}.New)
	r.DELETE("/product/:id", product.View{}.Delete)
}
