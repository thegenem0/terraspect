package service

import (
	"fmt"

	"github.com/thegenem0/terraspect_server/pkg/tree"
	"github.com/thegenem0/terraspect_server/repository"
)

type TreeService interface {
	BuildTree() (tree.TreeData, error)
}

type treeService struct {
	TreeRepository repository.TreeRepository
}

func NewTreeService(tr repository.TreeRepository) TreeService {
	return &treeService{
		TreeRepository: tr,
	}
}

func (ts *treeService) BuildTree() (tree.TreeData, error) {
	treeResult, err := ts.TreeRepository.BuildTree()
	if err != nil {
		return tree.TreeData{},
			fmt.Errorf("Failed to build tree: %s", err)

	}
	return treeResult, nil
}
