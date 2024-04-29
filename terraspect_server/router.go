package main

import (
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/handler"
	"github.com/thegenem0/terraspect_server/repository"
	"github.com/thegenem0/terraspect_server/service"
)

func InitRouter(modules *Modules) (*gin.Engine, error) {

	clerk, err := clerk.NewClient("sk_test_pEEsJt9JKgFJwcHtRYImJISLfqt92SOhLUksVv0g3N")
	if err != nil {
		return nil, err
	}

	authRepo := repository.NewAuthRepository(clerk)
	authService := service.NewAuthService(authRepo)

	router := gin.Default()

	baseURL := "/api/v1"

	handler.NewHandler(&handler.Config{
		R:               router,
		AuthService:     authService,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(5) * time.Second),
	})

	return router, nil
}
