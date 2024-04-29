package main

import (
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
	defer db.Close()

	return &Modules{
		DB: db,
	}, nil
}
