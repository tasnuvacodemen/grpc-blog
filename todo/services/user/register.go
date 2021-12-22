package user

import (
	"context"

	upb "github.com/tasnuvatina/grpc-blog/proto/user"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserSvc) Register(ctx context.Context, req *upb.RegisterUserRequest) (*upb.RegisterUserResponce, error) {
	user := storage.User{
		ID:       req.GetUser().ID,
		UserName: req.GetUser().UserName,
		Email:    req.GetUser().Email,
		Password: req.GetUser().Password,
	}

	id, err := s.core.Register(context.Background(), user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to register user")
	}
	return &upb.RegisterUserResponce{
		ID: id,
	}, nil
}
