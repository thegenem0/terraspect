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

	// if gin.Mode() != gin.TestMode {
	// 	g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
	// 	g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
	// 	g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
	// 	g.PUT("/details", middleware.AuthUser(h.TokenService), h.Details)
	// 	g.POST("/image", middleware.AuthUser(h.TokenService), h.Image)
	// 	g.DELETE("/image", middleware.AuthUser(h.TokenService), h.DeleteImage)
	// } else {

	g.OPTIONS("/tree", h.OptionsTree)
	g.GET("/tree", middleware.AuthMiddleware(c.AuthService), h.GetTree)

}
