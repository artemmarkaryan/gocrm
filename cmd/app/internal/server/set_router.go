package server

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/server/routes"
	"github.com/gin-gonic/gin"
)

type RouteGroup interface {
	AddRoutes(*gin.Engine)
}

var routeGroups = []RouteGroup{
	routes.Customer{},
	routes.Item{},
	routes.Order{},
	routes.Product{},
	routes.User{},
}

func setRouter(r *gin.Engine){
	for _, routeGroup := range routeGroups {
		routeGroup.AddRoutes(r)
	}
}