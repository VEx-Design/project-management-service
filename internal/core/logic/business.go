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
}

type UserService interface {
	GetUser(userId string) (*entities.User, error)
}
