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

	webGroup.OPTIONS("/apikey", h.OptionsAuth)
	webGroup.OPTIONS("/apikey/delete", h.OptionsAuth)

	webGroup.GET("/apikey", middleware.ClerkMiddleware(c.AuthService), h.GetAPIKeys)
	webGroup.POST("/apikey", middleware.ClerkMiddleware(c.AuthService), h.PostApiKey)
	webGroup.POST("/apikey/delete", middleware.ClerkMiddleware(c.AuthService), h.DeleteApiKey)

	webGroup.OPTIONS("/projects", h.OptionsProject)
	webGroup.GET("/projects", middleware.ClerkMiddleware(c.AuthService), h.GetAllProjects)
	webGroup.POST("/projects", middleware.ClerkMiddleware(c.AuthService), h.PostProject)

	webGroup.OPTIONS("/projects/:projectId", h.OptionsProject)
	webGroup.GET("/projects/:projectId", middleware.ClerkMiddleware(c.AuthService), h.GetProject)

	webGroup.OPTIONS("/projects/:projectId/plans", h.OptionsProject)
	webGroup.GET("/projects/:projectId/plans", middleware.ClerkMiddleware(c.AuthService), h.GetProjectPlans)

	webGroup.OPTIONS("/projects/:projectId/plans/:planId", h.OptionsProject)
	webGroup.GET("/projects/:projectId/plans/:planId", middleware.ClerkMiddleware(c.AuthService), h.GetProjectPlanById)

	webGroup.OPTIONS("/projects/:projectId/plans/:planId/graph", h.OptionsTree)
	webGroup.GET("/projects/:projectId/plans/:planId/graph", middleware.ClerkMiddleware(c.AuthService), h.GetTree)

	webGroup.OPTIONS("/projects/:projectId/plans/:planId/changes", h.OptionsTree)
	webGroup.GET("/projects/:projectId/plans/:planId/changes", middleware.ClerkMiddleware(c.AuthService), h.GetChanges)

	apiGroup.OPTIONS("/upload", h.OptionsProject)
	apiGroup.POST("/upload", middleware.ApiMiddleware(c.AuthService), h.PostProjectPlan)
}
