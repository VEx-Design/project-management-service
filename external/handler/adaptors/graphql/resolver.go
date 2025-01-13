package graph

import "project-management-service/internal/core/logic"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProjSrv logic.ProjectService
	UserSrv logic.UserService
}
