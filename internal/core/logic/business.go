package logic

import "project-management-service/internal/core/entities"

type ProjectService interface {
	CreateProject(project entities.Project) error
	// GetAllTypes() ([]entities.Type, error)
}
