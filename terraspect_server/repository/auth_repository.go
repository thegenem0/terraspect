package repository

import (
	"errors"
	"fmt"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/google/uuid"
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/pkg/database"
)

type AuthRepository interface {
	AuthenticateToken(token string) error
	GetUserData() *UserData
	GetUserByAPIKey(apiKey string) (string, error)
	GenerateAPIKey() (string, error)
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

type authRepository struct {
	clerk    clerk.Client
	db       database.IDBModule
	userData *UserData
}

func NewAuthRepository(clerk clerk.Client, db database.IDBModule) AuthRepository {
	return &authRepository{
		clerk:    clerk,
		db:       db,
		userData: &UserData{},
	}
}

func (ar *authRepository) getClerkSession(token string) (*clerk.SessionClaims, error) {
	session, err := ar.clerk.VerifyToken(token)

	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, errors.New("Invalid token")
	}

	return session, nil
}

func (ar *authRepository) getClerkUserData(claims jwt.Claims) (*UserData, error) {
	user, err := ar.clerk.Users().Read(claims.Subject)
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

func (ar *authRepository) AuthenticateToken(token string) error {
	session, err := ar.getClerkSession(token)
	if err != nil {
		return err
	}

	userData, err := ar.getClerkUserData(session.Claims)
	if err != nil {
		return err
	}

	ar.userData = userData
	return nil
}

func (ar *authRepository) GetUserData() *UserData {
	return ar.userData
}

func (ar *authRepository) GetUserByAPIKey(apiKey string) (string, error) {
	var apiKeyModel model.ApiKey
	result := ar.db.Connection().Where("key = ?", apiKey).First(&apiKeyModel)
	if result.Error != nil {
		return "", result.Error
	}

	var user model.User
	result = ar.db.Connection().Where("id = ?", apiKeyModel.UserID).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.ClerkUserID, nil
}

func (ar *authRepository) GenerateAPIKey() (string, error) {
	var user model.User
	result := ar.db.Connection().FirstOrCreate(&user, model.User{
		ClerkUserID: ar.userData.ID,
	})

	if result.Error != nil {
		fmt.Println(result.Error)
		return "", result.Error
	}

	apiKey := uuid.New().String()

	newApiKey := model.ApiKey{
		Key:    apiKey,
		UserID: user.ID,
	}

	if err := ar.db.Connection().Create(&newApiKey).Error; err != nil {
		fmt.Println(err)
		return "", err
	}

	return apiKey, nil
}
