package server

import (
	"github.com/artemmarkaryan/gocrm/cmd/app/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

type MiddlewareProvider interface {
	GetMiddlewareFunc() gin.HandlerFunc
}

var commonMiddlewareProviders = [...]MiddlewareProvider{
	middleware.SessionProvider{},
	middleware.CORSProvider{},
}

func setCommonMiddleware(r *gin.Engine) {
	for _, provider := range commonMiddlewareProviders {
		r.Use(provider.GetMiddlewareFunc())
	}
}
