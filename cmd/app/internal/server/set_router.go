package server

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/server/middleware"
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/server/routes"
	"github.com/gin-gonic/gin"
)

type RouteGroup interface {
	AddRoutes(group *gin.RouterGroup)
}

var privateRouteGroups = [...]RouteGroup{
	routes.Customer{},
	routes.Item{},
	routes.Order{},
	routes.Product{},
	routes.User{},
}

func setRouter(r *gin.Engine){
	routes.Auth{}.AddRoutes(r)

	private := r.Group("/")
	private.Use(middleware.AuthRequiredProvider{}.GetMiddlewareFunc())
	for _, routeGroup := range privateRouteGroups {
		routeGroup.AddRoutes(private)
	}
}