package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/handler/middleware"
	"github.com/thegenem0/terraspect_server/service"
)

type Handler struct {
	AuthService   service.AuthService
	TreeService   service.TreeService
	UploadService service.UploadService
}

type Config struct {
	R               *gin.Engine
	AuthService     service.AuthService
	TreeService     service.TreeService
	UploadService   service.UploadService
	WebBaseURL      string
	ApiBaseURL      string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	h := &Handler{
		TreeService:   c.TreeService,
		AuthService:   c.AuthService,
		UploadService: c.UploadService,
	}

	webGroup := c.R.Group(c.WebBaseURL)
	apiGroup := c.R.Group(c.ApiBaseURL)

	webGroup.OPTIONS("/apikey", h.OptionsAuth)
	webGroup.OPTIONS("/apikey/delete", h.OptionsAuth)

	webGroup.GET("/apikey", middleware.ClerkMiddleware(c.AuthService), h.GetAPIKeys)
	webGroup.POST("/apikey", middleware.ClerkMiddleware(c.AuthService), h.PostApiKey)
	webGroup.POST("/apikey/delete", middleware.ClerkMiddleware(c.AuthService), h.DeleteApiKey)

	webGroup.OPTIONS("/tree", h.OptionsTree)
	webGroup.GET("/tree", middleware.ClerkMiddleware(c.AuthService), h.GetTree)

	apiGroup.OPTIONS("/upload", h.OptionsUpload)
	apiGroup.POST("/upload", middleware.ApiMiddleware(c.AuthService), h.PostUpload)
}
