package repository

import (
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

func (r *projectRepositoryPQ) CreateProject(projectData entities.Project) error {
	newProject := gorm_model.Project{
		ID:          uuid.New().String(),
		OwnerId:     projectData.OwnerId,
		Name:        projectData.Name,
		Description: projectData.Description,
		Flow:        projectData.Flow,
	}

	var project gorm_model.Project
	if err := r.client.FirstOrCreate(&project, newProject).Error; err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
