package entities

import "time"

type Project struct {
	ID          string
	Name        string
	Description string
	Flow        string
	OwnerId     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
