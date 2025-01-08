package service

import (
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

func (s *projectService) GetMyProject(userId string) ([]entities.Project, error) {
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
