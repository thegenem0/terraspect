package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/thegenem0/terraspect_server/docs"
	"log"
	"os"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/handler"
	"github.com/thegenem0/terraspect_server/repository"
	"github.com/thegenem0/terraspect_server/service"
)

func InitRouter(modules *Modules) (*gin.Engine, error) {

	clerkApiKey := os.Getenv("CLERK_API_KEY")
	if clerkApiKey == "" {
		log.Panicf("CLERK_API_KEY is not set")
	}

	clerkClient, err := clerk.NewClient(clerkApiKey)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(modules.DB)
	projectRepo := repository.NewProjectRepository(modules.DB)

	authService := service.NewAuthService(clerkClient, userRepo)
	uploadService := service.NewUploadService(userRepo, projectRepo)
	treeService := service.NewTreeService(userRepo, projectRepo)
	projectService := service.NewProjectService(projectRepo)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	webBaseURL := "/api/web/v1"
	apiBaseUrl := "/api/v1"

	handler.NewHandler(&handler.Config{
		R:               router,
		AuthService:     authService,
		ProjectService:  projectService,
		TreeService:     treeService,
		UploadService:   uploadService,
		WebBaseURL:      webBaseURL,
		ApiBaseURL:      apiBaseUrl,
		TimeoutDuration: time.Duration(5) * time.Second,
	})

	return router, nil
}
