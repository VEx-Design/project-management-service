package ports

import "project-management-service/internal/core/entities"

type ProjectRepository interface {
	GetMyProject(userId string) ([]entities.Project, error)
	CreateProject(project entities.Project) error
	GetProject(projectId string) (*entities.Project, error)

	PublicShare(projectId string) error
	DepublicShare(projectId string) error
	GetPublicSharedProjects() ([]entities.Project, error)
	CanCloneProject(projectId string) (bool, error)
	CloneProject(projectId string, newOwnerId string) (*entities.Project, error)
	UpdateProject(project entities.Project, userId string) error
	DeleteProject(projectId string, userId string) error
}
