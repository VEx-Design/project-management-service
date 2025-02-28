package logic

import (
	"context"
	"project-management-service/internal/core/entities"
)

type ProjectService interface {
	CreateProject(project entities.Project) error
	GetMyProject(ctx context.Context, userId string) ([]entities.Project, error)
	GetProject(projectId string) (*entities.Project, error)
	UpdateProject(project entities.UpdateProject) error
	UpdateProjectFlow(project entities.UpdateProjectFlow) error
	DeleteProject(userid, projectid string) error

	PublicShare(projectId string) error
	GetPublicSharedProjects() ([]entities.Project, error)
	DepublicShare(projectId string) error
	CanCloneProject(projectId string) (bool, error)
	CloneProject(projectId string, newOwnerId string) (*entities.Project, error)
}

type UserService interface {
	GetUser(userId string) (*entities.User, error)
}
