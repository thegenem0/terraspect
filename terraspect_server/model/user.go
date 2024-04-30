package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ClerkUserID string
	ApiKeys     []ApiKey `gorm:"foreignKey:UserID"`
}

type ApiKey struct {
	gorm.Model
	UserID uint `gorm:"primaryKey"`
	Key    string
}
