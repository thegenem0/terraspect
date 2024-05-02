package dto

type ApiKeyResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Key         string `json:"key"`
	ProjectName string `json:"project_name"`
	CreatedAt   string `json:"created_at"`
}

type GenerateApiKeyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   string `json:"projectId"`
}

type DeleteApiKeyRequest struct {
	Key string `json:"key"`
}
