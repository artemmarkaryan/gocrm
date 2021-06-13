package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type CORSProvider struct{}

func (C CORSProvider) GetMiddlewareFunc() gin.HandlerFunc {
	log.Print("setting cors")
	return cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5000", "http://localhost", "http://0.0.0.0", "http://127.0.0.1", "https://golangcrm.herokuapp.com"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization",
			"accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

}
