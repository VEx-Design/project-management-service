package receiver

import (
	"context"

	ports "project-management-service/external/_ports"
	pb "project-management-service/external/receiver/adaptors/gRPC"
	"project-management-service/internal/core/entities"
)

type userReceiver struct {
	cilent pb.UserServiceClient
}

func NewUserReceiver(cilent pb.UserServiceClient) ports.UserReceiver {
	return &userReceiver{cilent: cilent}
}

func (r *userReceiver) GetUser(userId string) (*entities.User, error) {
	ctx := context.Background()
	req := &pb.GetUserRequest{UserId: userId}

	resp, err := r.cilent.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:      resp.User.Id,
		Name:    resp.User.Name,
		Email:   resp.User.Email,
		Picture: resp.User.Picture,
	}

	return user, nil
}
