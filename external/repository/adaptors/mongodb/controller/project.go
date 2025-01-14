package repository

import (
	"context"
	"fmt"
	"log"
	ports "project-management-service/external/_ports"
	mongo_model "project-management-service/external/repository/adaptors/mongodb/model"
	"project-management-service/internal/core/entities"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type projectRepositoryMongo struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewProjectRepositoryMongo(client *mongo.Client) ports.ProjectRepository {
	return &projectRepositoryMongo{
		client:   client,
		database: client.Database("project-management"),
	}
}

func (r *projectRepositoryMongo) CreateProject(project entities.Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	newProject := mongo_model.Project{
		OwnerId:     project.OwnerId,
		Name:        project.Name,
		Description: project.Description,
		Flow:        project.Flow,
	}

	insertResult, err := r.database.Collection("project").InsertOne(ctx, newProject)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)

	return nil
}

func (r *projectRepositoryMongo) GetMyProject(userId string) ([]entities.Project, error) {
	return nil, nil
}

func (r *projectRepositoryMongo) GetProject(projectId string) (*entities.Project, error) {
	return nil, nil
}
