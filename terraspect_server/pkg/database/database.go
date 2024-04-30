package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
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
	host := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName)

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
