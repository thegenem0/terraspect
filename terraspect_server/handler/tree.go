package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) OptionsTree(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (h *Handler) GetTree(c *gin.Context) {
	tree, _ := h.TreeService.BuildTree()

	c.JSON(200, gin.H{
		"tree": tree,
	})
}

func (h *Handler) PostTree(c *gin.Context) {

}
