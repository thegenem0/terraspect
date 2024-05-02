package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ClerkUserID string
	ApiKeys     []ApiKey `gorm:"foreignKey:UserID"`
	Plans       []Plan   `gorm:"foreignKey:UserID"`
}

type ApiKey struct {
	gorm.Model
	UserID      uint `gorm:"primaryKey"`
	Name        string
	Description string
	Key         string
}

type ApiKeyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Key         string `json:"key"`
	CreatedAt   string `json:"created_at"`
}
