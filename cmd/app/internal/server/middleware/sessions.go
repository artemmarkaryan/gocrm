package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
)

type SessionProvider struct{}

func (s SessionProvider) GetMiddlewareFunc() gin.HandlerFunc {
	log.Print("creating sessions")
	store := cookie.NewStore([]byte("secret"))

	return sessions.Sessions(
		"session",
		store,
	)
}
