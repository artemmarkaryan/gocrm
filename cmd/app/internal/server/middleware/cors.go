package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type CORSProvider struct {}

func (C CORSProvider) GetMiddlewareFunc() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization",
			"accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

}