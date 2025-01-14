package query

import (
	"context"
	"fmt"
	"project-management-service/external/handler/adaptors/graphql/model"
	"project-management-service/external/handler/router"
	"project-management-service/internal/core/logic"
	"time"
)

type ProjectQuery interface {
	GetProjects(ctx context.Context, ownerID string) ([]*model.Project, error)
	GetProject(ctx context.Context, id string) (*model.Project, error)
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

func (q projectQuery) GetProject(ctx context.Context, id string) (*model.Project, error) {
	userReq := ""
	userID := ctx.Value(router.RequestIDKey)
	if userID == nil {
		return nil, fmt.Errorf("user_id not found in context")
	}

	idStr, ok := userID.(string)
	if !ok {
		return nil, fmt.Errorf("user_id is not a valid string")
	}

	userReq = idStr // Assign the string value.

	proj, err := q.projSrv.GetProject(id)
	if err != nil {
		return nil, err
	}

	user, err := q.userSrv.GetUser(proj.OwnerId)
	if err != nil {
		return nil, err
	}

	if userReq != proj.OwnerId {
		return nil, fmt.Errorf("unauthorized access")
	}

	return &model.Project{
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
	}, nil
}
