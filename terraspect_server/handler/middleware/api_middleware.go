package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/model/apierror"
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
			apiErr := apierror.NewAPIError(
				apierror.InternalServerError,
				"Failed to bind header",
			)
			c.JSON(apiErr.Status(), gin.H{
				"error": apiErr,
			})

			c.Abort()
			return
		}

		clerkUserId, err := s.GetClerkUserIDFromAPIKey(h.ApiKey)
		if err != nil {
			apiErr := apierror.NewAPIError(
				apierror.APIKeyVerificationFailed,
				"API key verification failed",
			)
			c.JSON(apiErr.Status(), gin.H{
				"error": apiErr,
			})

			c.Abort()
			return
		}

		c.Set("clerkUserId", clerkUserId)

		c.Next()
	}
}
