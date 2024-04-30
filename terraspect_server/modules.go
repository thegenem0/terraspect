package main

import (
	"github.com/thegenem0/terraspect_server/model"
	"github.com/thegenem0/terraspect_server/pkg/change"
	"github.com/thegenem0/terraspect_server/pkg/database"
	"github.com/thegenem0/terraspect_server/pkg/reflector"
)

type Modules struct {
	DB        database.IDBModule
	Reflector reflector.IReflectorModule
	Change    change.IChangeModule
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
	})
	if err != nil {
		return nil, err
	}

	return &Modules{
		DB: db,
	}, nil
}
