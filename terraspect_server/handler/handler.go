package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/handler/middleware"
	"github.com/thegenem0/terraspect_server/service"
)

type Handler struct {
	AuthService    service.AuthService
	ProjectService service.ProjectService
	TreeService    service.TreeService
	UploadService  service.UploadService
}

type Config struct {
	R               *gin.Engine
	AuthService     service.AuthService
	ProjectService  service.ProjectService
	TreeService     service.TreeService
	UploadService   service.UploadService
	WebBaseURL      string
	ApiBaseURL      string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

func NewHandler(c *Config) {
	h := &Handler{
		AuthService:    c.AuthService,
		ProjectService: c.ProjectService,
		TreeService:    c.TreeService,
		UploadService:  c.UploadService,
	}

	webGroup := c.R.Group(c.WebBaseURL)
	apiGroup := c.R.Group(c.ApiBaseURL)

	c.R.GET("/health", h.GetHealth)

	webGroup.OPTIONS("/apikey/delete", h.OptionsAuth)

	webGroup.GET("/apikey", middleware.ClerkMiddleware(c.AuthService), h.GetAPIKeys)
	webGroup.POST("/apikey", middleware.ClerkMiddleware(c.AuthService), h.PostApiKey)
	webGroup.POST("/apikey/delete", middleware.ClerkMiddleware(c.AuthService), h.DeleteApiKey)

	webGroup.GET("/projects", middleware.ClerkMiddleware(c.AuthService), h.GetAllProjects)
	webGroup.POST("/projects", middleware.ClerkMiddleware(c.AuthService), h.PostProject)

	webGroup.GET("/projects/:projectId", middleware.ClerkMiddleware(c.AuthService), h.GetProject)
	webGroup.DELETE("/projects/:projectId", middleware.ClerkMiddleware(c.AuthService), h.DeleteProject)

	webGroup.GET("/projects/:projectId/plans", middleware.ClerkMiddleware(c.AuthService), h.GetProjectPlans)

	webGroup.GET("/projects/:projectId/plans/:planId", middleware.ClerkMiddleware(c.AuthService), h.GetProjectPlanById)

	webGroup.GET("/projects/:projectId/plans/:planId/graph", middleware.ClerkMiddleware(c.AuthService), h.GetTree)

	webGroup.GET("/projects/:projectId/plans/:planId/changes", middleware.ClerkMiddleware(c.AuthService), h.GetChanges)

	apiGroup.POST("/upload", middleware.ApiMiddleware(c.AuthService), h.PostProjectPlan)
}
