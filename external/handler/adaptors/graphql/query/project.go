package query

import (
	"project-management-service/external/handler/adaptors/graphql/model"
	"project-management-service/internal/core/logic"
)

type ProjectQuery interface {
	GetProjects(ownerID string) ([]*model.Project, error)
}

type projectQuery struct {
	projSrv logic.ProjectService
}

func NewProjectQuery(projectSrv logic.ProjectService) ProjectQuery {
	return &projectQuery{projSrv: projectSrv}
}

func (q projectQuery) GetProjects(ownerID string) ([]*model.Project, error) {
	projects, err := q.projSrv.GetMyProject(ownerID)
	if err != nil {
		return nil, err
	}

	var gqlProjects []*model.Project
	for _, proj := range projects {
		gqlProjects = append(gqlProjects, &model.Project{
			ID:          proj.ID,
			Name:        proj.Name,
			Description: &proj.Description,
			Flow:        &proj.Flow,
			OwnerID:     &proj.OwnerId,
		})
	}

	return gqlProjects, nil
}
