package user

import (
	"context"

	upb "github.com/tasnuvatina/grpc-blog/proto/user"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

type userCoreStore interface {
	Register(context.Context, storage.User) (int64, error)
	GetUser(context.Context, string) (*storage.User, error)
	GetUserById(context.Context, int64) (*storage.User, error)
}
type UserSvc struct {
	upb.UnimplementedTodoServiceServer
	core userCoreStore
}

func NewTodoServer(c userCoreStore) *UserSvc {
	return &UserSvc{
		core: c,
	}
}
