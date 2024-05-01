package service

import (
	"encoding/json"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/repository"
	"io"
	"mime/multipart"
)

type UploadService interface {
	SavePlanFile(
		clerkUserId string,
		file *multipart.FileHeader,
	) error
}

type uploadService struct {
	userRepository repository.UserRepository
}

func NewUploadService(ur repository.UserRepository) UploadService {
	return &uploadService{
		userRepository: ur,
	}
}

func (us *uploadService) SavePlanFile(
	clerkUserId string,
	file *multipart.FileHeader,
) error {
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

	err = us.userRepository.AddPlanToUser(clerkUserId, *planModel)
	if err != nil {
		return err
	}

	return nil
}
