package service

import (
	"context"
	"errors"
	"log"
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

func (s *projectService) GetMyProject(ctx context.Context, userId string) ([]entities.Project, error) {
	if userId == "" {
		return nil, errors.New("user ID is required")
	}

	projects, err := s.projRepo.GetMyProject(userId)
	if err != nil {
		log.Printf("failed to fetch projects for user %s: %v", userId, err)
		return nil, err
	}

	return projects, nil
}

func (s *projectService) GetProject(projectId string) (*entities.Project, error) {
	project, err := s.projRepo.GetProject(projectId)
	if err != nil {
		log.Printf("failed to fetch project %s: %v", projectId, err)
		return nil, err
	}

	return project, nil
}

func (s *projectService) UpdateProject(project entities.UpdateProject) error {
	error := s.projRepo.UpdateProject(project)
	if error != nil {
		log.Printf("failed to update project %s: %v", project.ID, error)
		return error
	}
	return nil
}

func (s *projectService) UpdateProjectFlow(project entities.UpdateProjectFlow) error {
	error := s.projRepo.UpdateProjectFlow(project)
	if error != nil {
		log.Printf("failed to update project %s flow: %v", project.ID, error)
		return error
	}
	return nil
}
