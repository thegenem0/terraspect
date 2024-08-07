package service

import (
	"errors"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/google/uuid"
	"github.com/thegenem0/terraspect_server/model/dto"
	"github.com/thegenem0/terraspect_server/repository"
)

type AuthService interface {
	GetUserID(token string) (string, error)
	GetClerkUserIDFromAPIKey(apiKey string) (string, error)
	GenerateAPIKey(request dto.GenerateApiKeyRequest) (string, error)
	GetAPIKeys() ([]dto.ApiKeyResponse, error)
	DeleteAPIKey(apiKey string) error
}

type UserData struct {
	ID                    string               `json:"id"`
	Username              *string              `json:"username"`
	FirstName             *string              `json:"first_name"`
	LastName              *string              `json:"last_name"`
	ProfileImageURL       string               `json:"profile_image_url"`
	PrimaryEmailAddressID *string              `json:"primary_email_address_id"`
	EmailAddresses        []clerk.EmailAddress `json:"email_addresses"`
}

type authService struct {
	clerkClient    clerk.Client
	userRepository repository.UserRepository
	userData       *UserData
}

func NewAuthService(
	clerkClient clerk.Client,
	userRepository repository.UserRepository,
) AuthService {
	return &authService{
		clerkClient:    clerkClient,
		userRepository: userRepository,
	}
}

func (as *authService) getClerkSession(token string) (*clerk.SessionClaims, error) {
	session, err := as.clerkClient.VerifyToken(token)

	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, errors.New("INVALID_TOKEN")
	}

	return session, nil
}

func (as *authService) getClerkUserData(claims jwt.Claims) (*UserData, error) {
	user, err := as.clerkClient.Users().Read(claims.Subject)
	if err != nil {
		return nil, err
	}

	userData := &UserData{
		ID:                    user.ID,
		Username:              user.Username,
		FirstName:             user.FirstName,
		LastName:              user.LastName,
		ProfileImageURL:       user.ProfileImageURL,
		PrimaryEmailAddressID: user.PrimaryEmailAddressID,
		EmailAddresses:        user.EmailAddresses,
	}
	return userData, nil
}

func (as *authService) GetUserID(token string) (string, error) {
	session, err := as.getClerkSession(token)
	if err != nil {
		return "", err
	}

	userData, err := as.getClerkUserData(session.Claims)
	if err != nil {
		return "", err
	}

	userExists, err := as.userRepository.CheckUserExists(userData.ID)
	if err != nil {
		return "", err
	}

	if !userExists {
		err = as.userRepository.CreateUser(userData.ID)
		if err != nil {
			return "", err
		}
	}

	as.userData = userData

	return userData.ID, nil
}

func (as *authService) GetClerkUserIDFromAPIKey(apiKey string) (string, error) {
	user, err := as.userRepository.GetUserFromAPIKey(apiKey)
	if err != nil {
		return "", err
	}

	return user.ClerkUserID, nil
}

func (as *authService) GetAPIKeys() ([]dto.ApiKeyResponse, error) {
	apiKeys, err := as.userRepository.GetAllAPIKeys(as.userData.ID)
	if err != nil {
		return nil, err
	}

	var keys []dto.ApiKeyResponse

	for _, apiKey := range apiKeys {
		apiKeyResponse := dto.ApiKeyResponse{
			ID:          apiKey.ID,
			Name:        apiKey.Name,
			Description: apiKey.Description,
			Key:         apiKey.Key,
			ProjectName: apiKey.Project.Name,
			CreatedAt:   apiKey.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		keys = append(keys, apiKeyResponse)
	}

	return keys, nil
}

func (as *authService) GenerateAPIKey(data dto.GenerateApiKeyRequest) (string, error) {
	apiKey := uuid.New().String()

	err := as.userRepository.AddAPIKey(as.userData.ID, apiKey, data)
	if err != nil {
		return "", err
	}

	return apiKey, nil
}

func (as *authService) DeleteAPIKey(apiKey string) error {
	return as.userRepository.DeleteAPIKey(apiKey)
}
