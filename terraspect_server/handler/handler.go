package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/handler/middleware"
	"github.com/thegenem0/terraspect_server/service"
)

type Handler struct {
	AuthService service.AuthService
	TreeService service.TreeService
}

type Config struct {
	R               *gin.Engine
	AuthService     service.AuthService
	TreeService     service.TreeService
	WebBaseURL      string
	ApiBaseURL      string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	h := &Handler{
		TreeService: c.TreeService,
		AuthService: c.AuthService,
	}

	webGroup := c.R.Group(c.WebBaseURL)
	_ = c.R.Group(c.ApiBaseURL)

	webGroup.OPTIONS("/apikey", h.OptionsAuth)
	webGroup.POST("/apikey", middleware.ClerkMiddleware(c.AuthService), h.PostApiKey)

	webGroup.OPTIONS("/tree", h.OptionsTree)
	webGroup.GET("/tree", middleware.ClerkMiddleware(c.AuthService), h.GetTree)

}
