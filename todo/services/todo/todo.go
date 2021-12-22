package todo

import (
	"context"

	tpb "github.com/tasnuvatina/grpc-blog/proto/todo"
	"github.com/tasnuvatina/grpc-blog/todo/storage"
)

type todoCoreStore interface {
	Create(context.Context, storage.Todo) (int64, error)
}
type Svc struct {
	tpb.UnimplementedTodoServiceServer
	core todoCoreStore
}

func NewTodoServer(c todoCoreStore) *Svc {
	return &Svc{
		core: c,
	}
}
