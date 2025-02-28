package gorm_model

import (
	"time"
)

type Project struct {
	ID              string `gorm:"primarykey"`
	Name            string `gorm:"not null"`
	OwnerId         string `gorm:"not null"`
	Description     string
	Flow            string
	ConfigurationID uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Shared          string
	SharedAccess    string //viewer , editor
	CloneAble       bool
}
