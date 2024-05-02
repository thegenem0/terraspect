package service

import (
	"github.com/thegenem0/terraspect_server/repository"
)

type UploadService interface {
}

type uploadService struct {
	userRepository    repository.UserRepository
	projectRepository repository.ProjectRepository
}

func NewUploadService(
	ur repository.UserRepository,
	pr repository.ProjectRepository,
) UploadService {
	return &uploadService{
		userRepository:    ur,
		projectRepository: pr,
	}
}
