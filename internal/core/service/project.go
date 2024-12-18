package service

import (
	ports "project-management-service/external/_ports"
	"project-management-service/internal/core/entities"
	"project-management-service/internal/core/logic"
)

type projectService struct {
	projRepo ports.ProjectRepository
}

func NewProjectService(projRepo ports.ProjectRepository) logic.ProjectService {
	return &projectService{projRepo: projRepo}
}

func (s *projectService) CreateProject(project entities.Project) error {
	s.projRepo.CreateProject(project)
	return nil
}
