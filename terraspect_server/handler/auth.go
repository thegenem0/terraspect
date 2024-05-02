package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/model/apierror"
	"github.com/thegenem0/terraspect_server/model/dto"
	"net/http"
)

type DeleteApiKeyBody struct {
	Key string `json:"key"`
}

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

func (h *Handler) GetAPIKeys(c *gin.Context) {
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

	keys, err := h.AuthService.GetAPIKeys()
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to get API keys",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"keys": keys,
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

	var body dto.GenerateApiKeyRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := apierror.NewAPIError(
			apierror.BadRequest,
			"Request body is invalid",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	apiKey, err := h.AuthService.GenerateAPIKey(body)
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

func (h *Handler) DeleteApiKey(c *gin.Context) {
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

	var body DeleteApiKeyBody
	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := apierror.NewAPIError(
			apierror.BadRequest,
			"Request body is invalid",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	err := h.AuthService.DeleteAPIKey(body.Key)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to delete API key",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
