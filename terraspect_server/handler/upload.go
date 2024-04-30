package handler

import "github.com/gin-gonic/gin"

func (h *Handler) OptionsUpload(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func (h *Handler) PostUpload(c *gin.Context) {

}
