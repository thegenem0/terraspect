package middleware

import (
	"github.com/thegenem0/terraspect_server/model/apierror"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/service"
)

type clerkAuthHeader struct {
	IDToken string `header:"Authorization"`
}

func ClerkMiddleware(s service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := clerkAuthHeader{}
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

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) < 2 {
			apiErr := apierror.NewAPIError(
				apierror.AuthorizationHeaderMissing,
				"Authorization header must be provided",
			)
			c.JSON(apiErr.Status(), gin.H{
				"error": apiErr,
			})

			c.Abort()
			return
		}

		clerkUserId, err := s.GetUserID(idTokenHeader[1])
		if err != nil {
			apiErr := apierror.NewAPIError(
				apierror.TokenVerificationFailed,
				"Token verification failed",
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
