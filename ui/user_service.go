package ui

import (
	"context"
	"github.com/go-leo/app-layout/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ api.UserService = (*UserService)(nil)

type UserService struct {
}

func (svc *UserService) CreateUser(ctx context.Context, request *api.CreateUserRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *UserService) ListUser(ctx context.Context, request *api.ListUserRequest) (*api.ListUserReply, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService() *UserService {
	return &UserService{}
}
