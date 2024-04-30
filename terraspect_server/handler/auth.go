package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) OptionsAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (h *Handler) PostApiKey(c *gin.Context) {
	apiKey, err := h.AuthService.GenerateAPIKey()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to generate API key",
		})
		return
	}

	c.JSON(200, gin.H{
		"key": apiKey,
	})
}
