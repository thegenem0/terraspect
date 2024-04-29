package repository

import (
	"errors"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

type AuthRepository interface {
	AuthenticateToken(token string) error
	GetUserData() *UserData
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
	userData *UserData
}

func NewAuthRepository(clerk clerk.Client) AuthRepository {
	return &authRepository{
		clerk:    clerk,
		userData: &UserData{},
	}
}

func (ar *authRepository) AuthenticateToken(token string) error {
	session, err := ar.clerk.VerifyToken(token)

	if err != nil {
		return err
	}

	if session == nil {
		return errors.New("Invalid token")
	}

	user, err := ar.clerk.Users().Read(session.Claims.Subject)
	if err != nil {
		return err
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

	ar.userData = userData
	return nil
}

func (ar *authRepository) GetUserData() *UserData {
	return ar.userData
}
