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

func (s *projectService) UpdateProject(project entities.Project, userId string) error {
	error := s.projRepo.UpdateProject(project, userId)
	if error != nil {
		log.Printf("failed to update project %s: %v", project.ID, error)
		return error
	}
	return nil
}

func (s *projectService) DeleteProject(projectId string, userId string) error {
	error := s.projRepo.DeleteProject(projectId, userId)
	if error != nil {
		log.Printf("failed to delete project %s: %v", projectId, error)
		return error
	}
	return nil
}

func (s *projectService) PublicShare(projectId string) error {
	error := s.projRepo.PublicShare(projectId)
	if error != nil {
		log.Printf("failed to share project %s: %v", projectId, error)
		return error
	}
	return nil
}

func (s *projectService) DepublicShare(projectId string) error {
	error := s.projRepo.DepublicShare(projectId)
	if error != nil {
		log.Printf("failed to unshare project %s: %v", projectId, error)
		return error
	}
	return nil
}

func (s *projectService) GetPublicSharedProjects() ([]entities.Project, error) {
	projects, err := s.projRepo.GetPublicSharedProjects()
	if err != nil {
		log.Printf("failed to fetch public shared projects: %v", err)
		return nil, err
	}

	return projects, nil
}

func (s *projectService) CanCloneProject(projectId string) (bool, error) {
	canClone, err := s.projRepo.CanCloneProject(projectId)
	if err != nil {
		log.Printf("failed to check if project %s can be cloned: %v", projectId, err)
		return false, err
	}

	return canClone, nil
}

func (s *projectService) CloneProject(projectId string, newOwnerId string) (*entities.Project, error) {
	shared, err := s.projRepo.CloneProject(projectId, newOwnerId)
	if err != nil {
		log.Printf("failed to clone project %s: %v", projectId, err)
		return nil, err
	}

	return shared, nil
}
