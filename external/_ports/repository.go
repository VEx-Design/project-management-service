package ports

import "project-management-service/internal/core/entities"

type ProjectRepository interface {
	// GetMyProject(userId string) error
	CreateProject(project entities.Project) error
}
