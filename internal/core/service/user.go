package service

import (
	ports "project-management-service/external/_ports"
	"project-management-service/internal/core/entities"
	"project-management-service/internal/core/logic"
)

type userService struct {
	userRev ports.UserReceiver
}

func NewUserService(userRev ports.UserReceiver) logic.UserService {
	return &userService{
		userRev: userRev,
	}
}

func (u *userService) GetUser(userId string) (*entities.User, error) {
	return u.userRev.GetUser(userId)
}
