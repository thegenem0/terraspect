package model

type User struct {
	Base
	ClerkUserID string
	ApiKeys     []ApiKey  `gorm:"foreignKey:UserID"`
	Projects    []Project `gorm:"foreignKey:UserID"`
}

type ApiKey struct {
	Base
	UserID      string `gorm:"index"`
	Name        string
	Description string
	Key         string
	ProjectID   string  `gorm:"index"`
	Project     Project `gorm:"foreignKey:ProjectID"`
}

type Project struct {
	Base
	UserID      string `gorm:"index"`
	Name        string
	Description string
	Plans       []Plan `gorm:"foreignKey:ProjectID"`
}

type Plan struct {
	Base
	ProjectID     string `gorm:"index"`
	TerraformPlan []byte
}
