package repository

import (
	"errors"
	"log"
	ports "project-management-service/external/_ports"
	gorm_model "project-management-service/external/repository/adaptors/postgres/model"
	"project-management-service/internal/core/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type projectRepositoryPQ struct {
	client *gorm.DB
}

func NewProjectRepositoryPQ(client *gorm.DB) ports.ProjectRepository {
	return &projectRepositoryPQ{
		client: client,
	}
}

// CreateProject creates a new project in the database.
func (r *projectRepositoryPQ) CreateProject(projectData entities.Project) error {
	if projectData.Name == "" || projectData.OwnerId == "" {
		return errors.New("invalid project data: name and owner ID are required")
	}

	newProject := gorm_model.Project{
		ID:              uuid.New().String(),
		OwnerId:         projectData.OwnerId,
		Name:            projectData.Name,
		Description:     projectData.Description,
		Flow:            projectData.Flow,
		ConfigurationID: projectData.ConfigurationID,
	}

	if err := r.client.Create(&newProject).Error; err != nil {
		log.Printf("failed to create project: %v", err)
		return err
	}

	return nil
}

// GetMyProject retrieves projects for the specified user.
func (r *projectRepositoryPQ) GetMyProject(userId string) ([]entities.Project, error) {
	if userId == "" {
		return nil, errors.New("user ID is required")
	}

	var projects []gorm_model.Project
	if err := r.client.Where("owner_id = ?", userId).Find(&projects).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entities.Project{}, nil
		}
		log.Printf("failed to retrieve projects for user %s: %v", userId, err)
		return nil, err
	}

	// Map database models to domain entities.
	var result []entities.Project
	for _, p := range projects {
		result = append(result, entities.Project{
			ID:              p.ID,
			OwnerId:         p.OwnerId,
			Name:            p.Name,
			Description:     p.Description,
			Flow:            p.Flow,
			ConfigurationID: p.ConfigurationID,
			CreatedAt:       p.CreatedAt,
			UpdatedAt:       p.UpdatedAt,
		})
	}

	return result, nil
}

// GetProject retrieves a project by its ID.
func (r *projectRepositoryPQ) GetProject(projectId string) (*entities.Project, error) {
	if projectId == "" {
		return nil, errors.New("project ID is required")
	}

	var project gorm_model.Project
	if err := r.client.Where("id = ?", projectId).First(&project).Error; err != nil {
		log.Printf("failed to retrieve project %s: %v", projectId, err)
		return nil, err
	}

	return &entities.Project{
		ID:              project.ID,
		OwnerId:         project.OwnerId,
		Name:            project.Name,
		Description:     project.Description,
		Flow:            project.Flow,
		ConfigurationID: project.ConfigurationID,
		CreatedAt:       project.CreatedAt,
		UpdatedAt:       project.UpdatedAt,
	}, nil
}

// update project flow
func (r *projectRepositoryPQ) UpdateProjectFlow(project entities.UpdateProjectFlow) error {
	if project.ID == "" || project.Flow == "" || project.UserID == "" {
		return errors.New("project ID, flow, and user ID are required")
	}

	if err := r.client.Model(&gorm_model.Project{}).Where("id = ? AND owner_id = ?", project.ID, project.UserID).Update("flow", project.Flow).Error; err != nil {
		log.Printf("failed to update project flow: %v", err)
		return err
	}

	return nil
}

// updates a project name and description in the database.
func (r *projectRepositoryPQ) UpdateProject(project entities.UpdateProject) error {
	if project.ID == "" || project.UserID == "" {
		return errors.New("project ID and user ID are required")
	}

	if err := r.client.Model(&gorm_model.Project{}).Where("id = ? AND owner_id = ?", project.ID, project.UserID).Updates(gorm_model.Project{
		Name:        project.Name,
		Description: project.Description,
	}).Error; err != nil {
		log.Printf("failed to update project: %v", err)
		return err
	}

	return nil
}
