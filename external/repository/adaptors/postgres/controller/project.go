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

func (r *projectRepositoryPQ) DeleteProject(userId string, projectId string) error {
	if userId == "" || projectId == "" {
		return errors.New("user ID and project ID are required")
	}

	if err := r.client.Where("id = ? AND owner_id = ?", projectId, userId).Delete(&gorm_model.Project{}).Error; err != nil {
		log.Printf("failed to delete project: %v", err)
		return err
	}

	return nil
}

// PublicShare updates the Shared column to 'public' and sets shared_access to 'view_only' for a project.
func (r *projectRepositoryPQ) PublicShare(projectId string) error {
	if projectId == "" {
		return errors.New("project ID is required")
	}

	// Find the project and update its Shared and SharedAccess columns
	result := r.client.Model(&gorm_model.Project{}).
		Where("id = ?", projectId).
		Updates(map[string]interface{}{
			"shared":        "public",
			"shared_access": "view_only",
		})

	if result.Error != nil {
		log.Printf("failed to share project publicly: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no project found with the given ID")
	}

	return nil
}

// DepublicShare updates the Shared column to 'no' and resets shared_access for a project.
func (r *projectRepositoryPQ) DepublicShare(projectId string) error {
	if projectId == "" {
		return errors.New("project ID is required")
	}

	// Find the project and update its Shared and SharedAccess columns
	result := r.client.Model(&gorm_model.Project{}).
		Where("id = ?", projectId).
		Updates(map[string]interface{}{
			"shared":        "",
			"shared_access": "",
		})

	if result.Error != nil {
		log.Printf("failed to remove public sharing: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no project found with the given ID")
	}

	return nil
}

// GetPublicSharedProjects retrieves all projects that are publicly shared.
func (r *projectRepositoryPQ) GetPublicSharedProjects() ([]entities.Project, error) {
	var projects []gorm_model.Project
	err := r.client.Where("shared = ?", "public").Find(&projects).Error
	if err != nil {
		log.Printf("failed to retrieve public shared projects: %v", err)
		return nil, err
	}

	// Map database models to domain entities
	result := make([]entities.Project, len(projects))
	for i, p := range projects {
		result[i] = entities.Project{
			ID:              p.ID,
			OwnerId:         p.OwnerId,
			Name:            p.Name,
			Description:     p.Description,
			Flow:            p.Flow,
			ConfigurationID: p.ConfigurationID,
			CreatedAt:       p.CreatedAt,
			UpdatedAt:       p.UpdatedAt,
			Shared:          p.Shared,
			SharedAccess:    p.SharedAccess,
			// CloneAble:       p.CloneAble,
		}
	}

	return result, nil
}

// CanCloneProject checks if a project is public and determines if it can be cloned.
func (r *projectRepositoryPQ) CanCloneProject(projectId string) (bool, error) {
	if projectId == "" {
		return false, errors.New("project ID is required")
	}

	// Find the project by ID and check if it's shared as 'public'
	var project gorm_model.Project
	err := r.client.Model(&gorm_model.Project{}).Where("id = ?", projectId).First(&project).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, errors.New("no project found with the given ID")
		}
		return false, err
	}

	// Check if the project is public
	if project.Shared == "public" {
		if err := r.client.Model(&project).Update("clone_able", true).Error; err != nil {
			log.Printf("failed to update CloneAble field for project %s: %v", projectId, err)
			return false, err
		}
		return true, nil
	}

	// If the project is not public, update the CloneAble field to false and return false
	if err := r.client.Model(&project).Update("clone_able", false).Error; err != nil {
		log.Printf("failed to update CloneAble field for project %s: %v", projectId, err)
		return false, err
	}
	return false, nil
}

// CloneProject clones an existing project by duplicating its details,
// assigning it a new owner, and setting the CloneAble field to true.
func (r *projectRepositoryPQ) CloneProject(projectId string, newOwnerId string) (*entities.Project, error) {
	if projectId == "" || newOwnerId == "" {
		return nil, errors.New("project ID and new owner ID are required")
	}

	// Find the project by ID
	var originalProject gorm_model.Project
	err := r.client.Model(&gorm_model.Project{}).Where("id = ?", projectId).First(&originalProject).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("no project found with the given ID")
		}
		return nil, err
	}

	// Check if the project is CloneAble
	if !originalProject.CloneAble {
		return nil, errors.New("project is not clonable")
	}

	// Create a new project entity by duplicating the original project data
	clonedProject := gorm_model.Project{
		Name:            originalProject.Name,
		Description:     originalProject.Description,
		OwnerId:         newOwnerId, // Set new owner
		Flow:            originalProject.Flow,
		ConfigurationID: originalProject.ConfigurationID,
		Shared:          originalProject.Shared,       // Retain the shared setting
		SharedAccess:    originalProject.SharedAccess, // Retain shared access
		CloneAble:       originalProject.CloneAble,    // Retain clone ability (which should already be true)
	}

	// Save the cloned project in the database
	if err := r.client.Create(&clonedProject).Error; err != nil {
		log.Printf("failed to clone project: %v", err)
		return nil, err
	}

	// Convert the cloned project into an entities.Project and return
	entitiesProject := &entities.Project{
		ID:              clonedProject.ID,
		Name:            clonedProject.Name,
		Description:     clonedProject.Description,
		OwnerId:         clonedProject.OwnerId,
		Flow:            clonedProject.Flow,
		ConfigurationID: clonedProject.ConfigurationID,
		Shared:          clonedProject.Shared,
		SharedAccess:    clonedProject.SharedAccess,
	}

	return entitiesProject, nil
}
