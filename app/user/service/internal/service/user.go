package service

import (
	"context"
	"mall-go/app/user/service/internal/biz"

	"google.golang.org/protobuf/types/known/emptypb"
	"mall-go/api/user/service"
)

type UserService struct {
	service.UnimplementedUserServer
	uu *biz.UserUsecase
}

func NewUserService(uu *biz.UserUsecase) *UserService {
	return &UserService{
		uu: uu,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *service.CreateUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *service.IdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) GetUserByAccount(ctx context.Context, req *service.AccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *service.IdsRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}