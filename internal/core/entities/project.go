package entities

import "time"

type Project struct {
	ID              string
	Name            string
	Description     string
	Flow            string
	TypesConfig     string
	OwnerId         string
	ConfigurationID uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Shared          string
	SharedAccess    string
	CloneAble       bool
}
