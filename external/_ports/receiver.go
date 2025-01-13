package ports

import "project-management-service/internal/core/entities"

type UserReceiver interface {
	GetUser(userId string) (*entities.User, error)
}
