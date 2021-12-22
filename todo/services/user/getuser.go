package user

import (
	"context"

	upb "github.com/tasnuvatina/grpc-blog/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserSvc) GetUser(ctx context.Context, req *upb.GetUserRequest) (*upb.GetUserResponce, error) {
	user, err := s.core.GetUser(context.Background(), req.UserName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to Get user")
	}
	return &upb.GetUserResponce{
		User: &upb.User{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}

func (s *UserSvc) GetUserById(ctx context.Context, req *upb.GetUserByIdRequest) (*upb.GetUserByIdResponce, error) {
	user, err := s.core.GetUserById(context.Background(), req.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to Get user by id")
	}
	return &upb.GetUserByIdResponce{
		User: &upb.User{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email,
			Password: user.Password,
		},
	}, nil
}
