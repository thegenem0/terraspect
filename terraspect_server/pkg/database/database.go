package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDBModule interface {
	Close() error
	SyncDatabase(models []interface{}) error
	Connection() *gorm.DB
}

type DBModule struct {
	connection *gorm.DB
}

func NewDBModule() (*DBModule, error) {
	connString := "host=localhost user=terraspect_root password=SuperSecretPassword dbname=terraspect_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DBModule{
		connection: db,
	}, nil
}

func (dbs *DBModule) Close() error {
	sqlDB, err := dbs.connection.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (dbs *DBModule) SyncDatabase(models []interface{}) error {
	for _, model := range models {
		err := dbs.connection.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dbs *DBModule) Connection() *gorm.DB {
	return dbs.connection
}
