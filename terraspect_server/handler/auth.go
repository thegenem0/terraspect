package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/model/apierror"
	"net/http"
)

// OptionsAuth godoc
// @Summary Options for auth
// @Schemes
// @Description Options for auth
// @Tags Auth
// @Success 200
// @Router /web/v1/apikey [options]
func (h *Handler) OptionsAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

// PostApiKey godoc
// @Summary Generate an API key
// @Schemes
// @Description Generate an API key
// @Tags Auth
// @securityDefinitions.apiKey ApiKeyAuth
// @Success 200
// @Router /web/v1/apikey [post]
func (h *Handler) PostApiKey(c *gin.Context) {
	_, exists := c.Get("clerkUserId")
	if !exists {
		apiErr := apierror.NewAPIError(
			apierror.TokenVerificationFailed,
			"User has no valid session",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	apiKey, err := h.AuthService.GenerateAPIKey()
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to generate API key",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"key": apiKey,
	})
}
