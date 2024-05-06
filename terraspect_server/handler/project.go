package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thegenem0/terraspect_server/model/apierror"
	"github.com/thegenem0/terraspect_server/model/dto"
	"net/http"
)

func (h *Handler) OptionsProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (h *Handler) GetProject(c *gin.Context) {
	projectId := c.Param("projectId")
	project, err := h.ProjectService.GetProjectByID(projectId)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			err.Error(),
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project": project,
	})
}

func (h *Handler) GetAllProjects(c *gin.Context) {
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

	projects, err := h.ProjectService.GetAllProjectsByUser(clerkUserId.(string))
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to get projects",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
	})
}

func (h *Handler) PostProject(c *gin.Context) {
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

	var body dto.PostProjectApiRequest
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

	err := h.ProjectService.CreateProject(clerkUserId.(string), body)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			err.Error(),
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project created",
	})
}

func (h *Handler) DeleteProject(c *gin.Context) {
	projectId := c.Param("projectId")

	err := h.ProjectService.DeleteProject(projectId)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to delete project",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project deleted",
	})
}

func (h *Handler) GetProjectPlans(c *gin.Context) {
	projectId := c.Param("projectId")

	plans, err := h.ProjectService.GetAllPlansByProject(projectId)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to get plans",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"plans": plans,
	})
}

func (h *Handler) GetProjectPlanById(c *gin.Context) {
	projectId := c.Param("projectId")
	planId := c.Param("planId")

	plan, err := h.ProjectService.GetPlanByID(projectId, planId)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to get plan",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"plan": plan,
	})
}

func (h *Handler) PostProjectPlan(c *gin.Context) {
	apiKey, exists := c.Get("apiKey")
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

	err = h.ProjectService.AddPlanToProject(apiKey.(string), file)
	if err != nil {
		apiErr := apierror.NewAPIError(
			apierror.InternalServerError,
			"Failed to create plan",
		)
		c.JSON(apiErr.Status(), gin.H{
			"error": apiErr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Plan created",
	})
}
