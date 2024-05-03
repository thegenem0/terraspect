package main

import (
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/pkg/database"
)

type Modules struct {
	DB database.IDBModule
}

func initModules() (*Modules, error) {
	db, err := database.NewDBModule()
	if err != nil {
		return nil, err
	}

	err = db.SyncDatabase([]interface{}{
		model.Plan{},
		model.User{},
		model.ApiKey{},
		model.Project{},
	})
	if err != nil {
		return nil, err
	}

	return &Modules{
		DB: db,
	}, nil
}
