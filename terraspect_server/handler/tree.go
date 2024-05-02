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
	projectId := c.Param("projectId")
	planId := c.Param("planId")

	tree, err := h.TreeService.BuildTree(projectId, planId)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.BadRequest,
			"Incorrect project ID provided",
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
