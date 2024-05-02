package dto

type GetProjectApiRequest struct {
	ProjectID string `uri:"project_id" binding:"required"`
}

type GetPlanApiRequest struct {
	PlanID string `uri:"plan_id" binding:"required"`
}

type PostProjectApiRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type GetAllProjectsResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PlanCount   int    `json:"planCount"`
	CreatedAt   string `json:"createdAt"`
}

type GetProjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Plans       []PlanApiResponse
	CreatedAt   string `json:"createdAt"`
}
