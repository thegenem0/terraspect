package service

import (
	"encoding/json"
	"fmt"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/thegenem0/terraspect_server/pkg/tree"
	"github.com/thegenem0/terraspect_server/repository"
)

type TreeService interface {
	BuildTree(clerkUserId string) (tree.TreeData, error)
}

type treeService struct {
	userRepository repository.UserRepository
}

func NewTreeService(ur repository.UserRepository) TreeService {
	return &treeService{
		userRepository: ur,
	}
}

func (ts *treeService) BuildTree(clerkUserId string) (tree.TreeData, error) {
	plan, err := ts.userRepository.GetLatestPlan(clerkUserId)
	if err != nil {
		return tree.TreeData{}, err
	}

	var storedPlan *tfjson.Plan

	err = json.Unmarshal(plan.TerraformPlan, &storedPlan)
	if err != nil {
		return tree.TreeData{},
			fmt.Errorf("Failed to unmarshal plan: %s", err)
	}

	treeData, err := tree.BuildTree(storedPlan.PlannedValues.RootModule)
	if err != nil {
		return tree.TreeData{}, err
	}

	return treeData, nil
}
