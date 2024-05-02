package repository

import (
	"fmt"
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/model/dto"
	"github.com/thegenem0/terraspect_server/pkg/database"
)

type UserRepository interface {
	GetUserFromAPIKey(apiKey string) (model.User, error)
	AddAPIKey(clerkUserID string, apiKey string, data dto.GenerateApiKeyRequest) error
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

func (ur *userRepository) GetUserFromAPIKey(apiKey string) (model.User, error) {
	var apiKeyModel model.ApiKey
	result := ur.db.Connection().Where("key = ?", apiKey).First(&apiKeyModel)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	var user model.User
	result = ur.db.Connection().Where("id = ?", apiKeyModel.UserID).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	if user.ClerkUserID == "" {
		return model.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (ur *userRepository) AddAPIKey(clerkUserID string, apiKey string, data dto.GenerateApiKeyRequest) error {
	var user model.User
	result := ur.db.Connection().First(&user, "clerk_user_id = ?", clerkUserID)
	if result.Error != nil {
		return result.Error
	}

	var project model.Project
	result = ur.db.Connection().First(&project, "id = ?", data.ProjectId)
	if result.Error != nil {
		return result.Error
	}

	newApiKey := model.ApiKey{
		Name:        data.Name,
		Description: data.Description,
		Key:         apiKey,
		Project:     project,
		UserID:      user.ID,
	}

	if err := ur.db.Connection().Create(&newApiKey).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
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
