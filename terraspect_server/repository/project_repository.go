package repository

import (
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/pkg/database"
)

type ProjectRepository interface {
	CreateProject(clerkUserID string, project model.Project) error
	GetProjectByID(projectID string) (model.Project, error)
	GetProjectByAPIKey(apiKey string) (model.Project, error)
	GetAllProjectsByUser(clerkUserId string) ([]model.Project, error)
	UpdateProject(project model.Project) error
	DeleteProject(projectID string) error
	GetAllPlansByProject(projectID string) ([]model.Plan, error)
	GetPlanByID(projectId string, planID string) (model.Plan, error)
	AddPlanToProject(projectID string, plan model.Plan) error
}

type projectRepository struct {
	db database.IDBModule
}

func NewProjectRepository(
	db database.IDBModule,
) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (pr *projectRepository) CreateProject(clerkUserID string, project model.Project) error {
	var user model.User
	result := pr.db.Connection().Where("clerk_user_id = ?", clerkUserID).First(&user)
	if result.Error != nil {
		return result.Error
	}

	project.UserID = user.ID
	result = pr.db.Connection().Create(&project)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *projectRepository) GetProjectByID(projectID string) (model.Project, error) {
	var project model.Project
	result := pr.db.Connection().Where("id = ?", projectID).First(&project)
	if result.Error != nil {
		return model.Project{}, result.Error
	}

	return project, nil
}

func (pr *projectRepository) GetProjectByAPIKey(apiKey string) (model.Project, error) {
	var apiKeyModel model.ApiKey
	result := pr.db.Connection().Where("key = ?", apiKey).First(&apiKeyModel)
	if result.Error != nil {
		return model.Project{}, result.Error
	}

	var project model.Project
	result = pr.db.Connection().Where("id = ?", apiKeyModel.ProjectID).First(&project)
	if result.Error != nil {
		return model.Project{}, result.Error
	}

	return project, nil
}

func (pr *projectRepository) GetAllProjectsByUser(clerkUserId string) ([]model.Project, error) {
	var user model.User
	result := pr.db.Connection().Where("clerk_user_id = ?", clerkUserId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	var projects []model.Project
	result = pr.db.Connection().Where("user_id = ?", user.ID).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}

	return projects, nil
}

func (pr *projectRepository) UpdateProject(project model.Project) error {
	result := pr.db.Connection().Save(&project)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *projectRepository) DeleteProject(projectID string) error {
	project := model.Project{
		Base: model.Base{
			ID: projectID,
		},
	}

	plans := pr.db.Connection().Where("project_id = ?", projectID).Delete(&model.Plan{})
	if plans.Error != nil {
		return plans.Error
	}

	apiKeys := pr.db.Connection().Where("project_id = ?", projectID).Delete(&model.ApiKey{})
	if apiKeys.Error != nil {
		return apiKeys.Error
	}

	result := pr.db.Connection().Delete(&project)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *projectRepository) GetAllPlansByProject(projectID string) ([]model.Plan, error) {
	var project model.Project
	result := pr.db.Connection().Where("id = ?", projectID).First(&project)
	if result.Error != nil {
		return nil, result.Error
	}

	var plans []model.Plan
	result = pr.db.Connection().Where("project_id = ?", project.ID).Find(&plans)
	if result.Error != nil {
		return nil, result.Error
	}

	return plans, nil
}

func (pr *projectRepository) GetPlanByID(projectId string, planID string) (model.Plan, error) {
	var plan model.Plan
	result := pr.db.Connection().Where("id = ? AND project_id = ?", planID, projectId).First(&plan)
	if result.Error != nil {
		return model.Plan{}, result.Error
	}

	return plan, nil
}

func (pr *projectRepository) AddPlanToProject(projectID string, plan model.Plan) error {
	var project model.Project
	result := pr.db.Connection().Where("id = ?", projectID).First(&project)
	if result.Error != nil {
		return result.Error
	}

	plan.ProjectID = project.ID
	result = pr.db.Connection().Create(&plan)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *projectRepository) DeletePlan(planID string) error {
	plan := model.Plan{
		Base: model.Base{
			ID: planID,
		},
	}
	result := pr.db.Connection().Delete(&plan)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
