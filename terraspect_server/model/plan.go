package model

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	UserID        uint `gorm:"primaryKey"`
	TerraformPlan []byte
}
