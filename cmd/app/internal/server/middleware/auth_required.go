package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const UserKey = "userId"

type AuthRequiredProvider struct{}

func (a AuthRequiredProvider) GetMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("checking auth")
		session := sessions.Default(c)
		user := session.Get(UserKey)
		if user == nil {
			// Abort the request with the appropriate error code
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Continue down the chain to handler etc
		c.Next()
	}
}
