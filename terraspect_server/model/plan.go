package model

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	ID            uint
	TerraformPlan []byte
}
