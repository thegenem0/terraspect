package service

import (
	"fmt"

	"github.com/thegenem0/terraspect_server/repository"
)

type AuthService interface {
	VerifyToken(token string) (bool, error)
	GetUserID(token string) (string, error)
}

type authService struct {
	AuthRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		AuthRepository: authRepository,
	}
}

func (as *authService) VerifyToken(token string) (bool, error) {
	as.AuthRepository.AuthenticateToken(token)
	return as.AuthRepository.GetUserData().ID != "", nil
}

func (as *authService) GetUserID(token string) (string, error) {
	if as.AuthRepository.GetUserData().ID == "" {
		return "", fmt.Errorf("User not found")
	}
	return as.AuthRepository.GetUserData().ID, nil
}
