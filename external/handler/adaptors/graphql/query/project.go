package query

import (
	"context"
	"fmt"
	"project-management-service/external/handler/adaptors/graphql/model"
	"project-management-service/internal/core/logic"
	"time"
)

type ProjectQuery interface {
	GetProjects(ctx context.Context, ownerID string) ([]*model.Project, error)
}

type projectQuery struct {
	projSrv logic.ProjectService
}

func NewProjectQuery(projectSrv logic.ProjectService) ProjectQuery {
	return &projectQuery{projSrv: projectSrv}
}

func (q projectQuery) GetProjects(ctx context.Context, ownerID string) ([]*model.Project, error) {
	projects, err := q.projSrv.GetMyProject(ctx, ownerID)
	fmt.Println("ownerID", ownerID)
	if err != nil {
		return nil, err
	}

	gqlProjects := make([]*model.Project, len(projects))
	for i, proj := range projects {
		gqlProjects[i] = &model.Project{
			ID:          proj.ID,
			Name:        proj.Name,
			Description: &proj.Description,
			Flow:        &proj.Flow,
			OwnerID:     &proj.OwnerId,
			CreatedAt:   proj.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   proj.UpdatedAt.Format(time.RFC3339),
		}
	}

	return gqlProjects, nil
}
