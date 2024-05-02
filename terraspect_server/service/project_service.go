package service

import (
	"encoding/json"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/model/dto"
	"github.com/thegenem0/terraspect_server/repository"
	"io"
	"mime/multipart"
)

type ProjectService interface {
	CreateProject(userID string, project dto.PostProjectApiRequest) error
	GetProjectByID(projectID string) (dto.GetProjectResponse, error)
	GetAllProjectsByUser(userID string) ([]dto.GetAllProjectsResponse, error)
	DeleteProject(projectID string) error
	GetAllPlansByProject(projectID string) ([]dto.PlanApiResponse, error)
	GetPlanByID(projectId string, planId string) (dto.PlanApiResponse, error)
	AddPlanToProject(projectID string, file *multipart.FileHeader) error
}

type projectService struct {
	projectRepository repository.ProjectRepository
}

func NewProjectService(
	pr repository.ProjectRepository,
) ProjectService {
	return &projectService{
		projectRepository: pr,
	}
}

func (ps *projectService) CreateProject(clerkUserId string, project dto.PostProjectApiRequest) error {
	projectModel := model.Project{
		Name:        project.Name,
		Description: project.Description,
	}
	return ps.projectRepository.CreateProject(clerkUserId, projectModel)
}

func (ps *projectService) GetProjectByID(projectID string) (dto.GetProjectResponse, error) {
	project, err := ps.projectRepository.GetProjectByID(projectID)
	if err != nil {
		return dto.GetProjectResponse{}, err
	}

	var projectPlans []dto.PlanApiResponse
	for _, plan := range project.Plans {
		projectPlans = append(projectPlans, dto.PlanApiResponse{
			ID:        plan.ID,
			ProjectID: plan.ProjectID,
			CreatedAt: plan.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response := dto.GetProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		Plans:       projectPlans,
		CreatedAt:   project.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (ps *projectService) GetAllProjectsByUser(clerkUserId string) ([]dto.GetAllProjectsResponse, error) {
	projects, err := ps.projectRepository.GetAllProjectsByUser(clerkUserId)
	if err != nil {
		return nil, err
	}

	var projectResponse []dto.GetAllProjectsResponse
	for _, project := range projects {
		projectResponse = append(projectResponse, dto.GetAllProjectsResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			PlanCount:   len(project.Plans),
			CreatedAt:   project.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return projectResponse, nil
}

func (ps *projectService) DeleteProject(projectID string) error {
	return ps.projectRepository.DeleteProject(projectID)
}

func (ps *projectService) GetAllPlansByProject(projectID string) ([]dto.PlanApiResponse, error) {
	plans, err := ps.projectRepository.GetAllPlansByProject(projectID)
	if err != nil {
		return []dto.PlanApiResponse{}, err
	}

	var planResponse []dto.PlanApiResponse
	for _, plan := range plans {
		planResponse = append(planResponse, dto.PlanApiResponse{
			ID:        plan.ID,
			ProjectID: plan.ProjectID,
			CreatedAt: plan.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return planResponse, nil
}

func (ps *projectService) GetPlanByID(projectId string, planId string) (dto.PlanApiResponse, error) {
	plan, err := ps.projectRepository.GetPlanByID(projectId, planId)
	if err != nil {
		return dto.PlanApiResponse{}, err
	}

	response := dto.PlanApiResponse{
		ID:        plan.ID,
		ProjectID: plan.ProjectID,
		CreatedAt: plan.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return response, nil
}

func (ps *projectService) AddPlanToProject(apiKey string, file *multipart.FileHeader) error {
	project, err := ps.projectRepository.GetProjectByAPIKey(apiKey)
	if err != nil {
		return err
	}

	contents, err := file.Open()
	if err != nil {
		return err
	}
	defer func(contents multipart.File) {
		err := contents.Close()
		if err != nil {
			return
		}
	}(contents)

	fileData, err := io.ReadAll(contents)
	if err != nil {
		return err
	}

	var plan *tfjson.Plan
	err = json.Unmarshal(fileData, &plan)
	if err != nil {
		return err
	}

	planModel := &model.Plan{
		TerraformPlan: fileData,
	}
	return ps.projectRepository.AddPlanToProject(project.ID, *planModel)
}
