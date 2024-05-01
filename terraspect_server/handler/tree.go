package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/model/apierror"
	"net/http"
)

func (h *Handler) OptionsTree(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (h *Handler) GetTree(c *gin.Context) {
	clerkUserId, exists := c.Get("clerkUserId")
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

	tree, err := h.TreeService.BuildTree(clerkUserId.(string))
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to build tree",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tree": tree,
	})
}

func (h *Handler) PostTree(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
