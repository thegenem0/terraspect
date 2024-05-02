package repository

import (
	"fmt"
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/pkg/database"
)

type UserRepository interface {
	GetClerkIDFromAPIKey(apiKey string) (string, error)
	AddAPIKeyToUser(clerkUserID string, apiKey string, name string, description string) error
	AddPlanToUser(clerkUserID string, plan model.Plan) error
	GetLatestPlan(clerkUserID string) (model.Plan, error)
	GetAllAPIKeys(clerkUserID string) ([]model.ApiKey, error)
	DeleteAPIKey(apiKey string) error
}

type userRepository struct {
	db database.IDBModule
}

func NewUserRepository(
	db database.IDBModule,
) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetClerkIDFromAPIKey(apiKey string) (string, error) {
	var apiKeyModel model.ApiKey
	result := ur.db.Connection().Where("key = ?", apiKey).First(&apiKeyModel)
	if result.Error != nil {
		return "", result.Error
	}

	var user model.User
	result = ur.db.Connection().Where("id = ?", apiKeyModel.UserID).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	if user.ClerkUserID == "" {
		return "", result.Error
	}

	return user.ClerkUserID, nil
}

func (ur *userRepository) AddAPIKeyToUser(clerkUserID string, apiKey string, name string, description string) error {
	var user model.User
	result := ur.db.Connection().FirstOrCreate(&user, model.User{
		ClerkUserID: clerkUserID,
	})

	if result.Error != nil {
		return result.Error
	}

	newApiKey := model.ApiKey{
		Name:        name,
		Description: description,
		Key:         apiKey,
		UserID:      user.ID,
	}

	if err := ur.db.Connection().Create(&newApiKey).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (ur *userRepository) AddPlanToUser(clerkUserID string, plan model.Plan) error {
	var user model.User
	result := ur.db.Connection().Where("clerk_user_id = ?", clerkUserID).First(&user)
	if result.Error != nil {
		return result.Error
	}

	plan.UserID = user.ID

	if err := ur.db.Connection().Create(&plan).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetLatestPlan(clerkUserID string) (model.Plan, error) {
	var user model.User
	result := ur.db.Connection().Where("clerk_user_id = ?", clerkUserID).First(&user)
	if result.Error != nil {
		return model.Plan{}, result.Error
	}

	var plan model.Plan
	result = ur.db.Connection().Where("user_id = ?", user.ID).Last(&plan)
	if result.Error != nil {
		return model.Plan{}, result.Error
	}

	return plan, nil
}

func (ur *userRepository) GetAllAPIKeys(clerkUserID string) ([]model.ApiKey, error) {
	var user model.User
	result := ur.db.Connection().Where("clerk_user_id = ?", clerkUserID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	var apiKeys []model.ApiKey
	result = ur.db.Connection().Where("user_id = ? AND deleted_at IS NULL", user.ID).Find(&apiKeys)
	if result.Error != nil {
		return nil, result.Error
	}

	return apiKeys, nil
}

func (ur *userRepository) DeleteAPIKey(apiKey string) error {
	if err := ur.db.Connection().
		Model(&model.ApiKey{}).
		Where("key = ?", apiKey).
		Update("deleted_at", "NOW()").
		Error; err != nil {
		return err
	}

	return nil
}
