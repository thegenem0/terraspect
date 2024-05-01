package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/model/apierror"
	"net/http"
)

func (h *Handler) OptionsUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (h *Handler) PostUpload(c *gin.Context) {
	clerkUserId, exists := c.Get("clerkUserId")
	if !exists {
		apiErr := apierror.NewAPIError(
			apierror.APIKeyVerificationFailed,
			"User has no valid session",
		)

		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	file, err := c.FormFile("plan_file")
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.NoFile,
			"Failed to get file",
		)
		c.JSON(apiErr.Status(), gin.H{
			"message": apiErr,
		})
		return
	}
	projectName := c.PostForm("project_name")
	if projectName == "" {
		apiErr := apierror.NewAPIError(
			apierror.NoProjectName,
			"No project name provided",
		)

		c.JSON(apiErr.Status(), gin.H{
			"message": apiErr,
		})
		return
	}

	err = h.UploadService.SavePlanFile(clerkUserId.(string), file)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to save file",
		)
		c.JSON(apiErr.Status(), gin.H{
			"message": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload processed",
	})
}
