package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/service"
)

type apiAuthHeader struct {
	ApiKey string `header:"X-Api-Key"`
}

func ApiMiddleware(s service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := apiAuthHeader{}
		err := c.ShouldBindHeader(&h)
		if err != nil {
			return
		}

		clerkUserId, err := s.GetUserByAPIKey(h.ApiKey)
		if err != nil {
			c.JSON(401, gin.H{
				"error": "API key verification failed",
			})

			c.Abort()
			return
		}
		fmt.Println("CLERK_USER_ID: ", clerkUserId)

		c.Next()
	}
}
