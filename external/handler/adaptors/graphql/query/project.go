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
	userSrv logic.UserService
}

func NewProjectQuery(projectSrv logic.ProjectService, userSrv logic.UserService) ProjectQuery {
	return &projectQuery{projSrv: projectSrv, userSrv: userSrv}
}

func (q projectQuery) GetProjects(ctx context.Context, ownerID string) ([]*model.Project, error) {
	projects, err := q.projSrv.GetMyProject(ctx, ownerID)
	fmt.Println("ownerID", ownerID)
	if err != nil {
		return nil, err
	}

	gqlProjects := make([]*model.Project, len(projects))
	for i, proj := range projects {
		user, err := q.userSrv.GetUser(proj.OwnerId)
		if err != nil {
			return nil, err
		}

		gqlProjects[i] = &model.Project{
			ID:          proj.ID,
			Name:        proj.Name,
			Description: &proj.Description,
			Flow:        &proj.Flow,
			Owner: &model.User{
				Name:    user.Name,
				Picture: &user.Picture,
			},
			CreatedAt: proj.CreatedAt.Format(time.RFC3339),
			UpdatedAt: proj.UpdatedAt.Format(time.RFC3339),
		}
	}

	return gqlProjects, nil
}
