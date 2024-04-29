package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/handler/middleware"
	"github.com/thegenem0/terraspect_server/service"
)

type Handler struct {
	AUthService service.AuthService
	TreeService service.TreeService
}

type Config struct {
	R               *gin.Engine
	AuthService     service.AuthService
	TreeService     service.TreeService
	BaseURL         string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	h := &Handler{
		TreeService: c.TreeService,
	}

	g := c.R.Group(c.BaseURL)

	g.OPTIONS("/tree", h.OptionsTree)
	g.GET("/tree", middleware.AuthMiddleware(c.AuthService), h.GetTree)

}
