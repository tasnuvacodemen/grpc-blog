package todo

import (
	"context"

	tpb "github.com/tasnuvatina/grpc-blog/proto/todo"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) Create(ctx context.Context, req *tpb.CreateTodoRequest) (*tpb.CreateTodoResponse, error) {
	// need to validate request
	todo := storage.Todo{
		ID:          req.GetTodo().ID,
		Title:       req.GetTodo().Title,
		Description: req.GetTodo().Description,
	}
	id, err := s.core.Create(context.Background(), todo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create todo")
	}
	return &tpb.CreateTodoResponse{
		ID: id,
	}, nil
}
