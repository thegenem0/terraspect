package repository

import (
	"encoding/json"
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/pkg/database"
	"github.com/thegenem0/terraspect_server/pkg/tree"
)

type TreeRepository interface {
	BuildTree() (tree.TreeData, error)
}

type treeRepository struct {
	Database database.IDBModule
}

func NewTreeRepository(
	database database.IDBModule,
) TreeRepository {
	return &treeRepository{
		Database: database,
	}
}

func (tr *treeRepository) BuildTree() (tree.TreeData, error) {
	var plan model.Plan
	tr.Database.Connection().First(&plan)

	var storedPlan *tfjson.Plan

	err := json.Unmarshal(plan.TerraformPlan, &storedPlan)
	if err != nil {
		return tree.TreeData{},
			fmt.Errorf("Failed to unmarshal plan: %s", err)
	}

	fmt.Println(storedPlan.PlannedValues.RootModule)

	treeResult, err := tree.BuildTree(storedPlan.PlannedValues.RootModule)
	if err != nil {
		return tree.TreeData{},
			fmt.Errorf("Failed to build tree: %s", err)
	}

	return treeResult, nil
}
